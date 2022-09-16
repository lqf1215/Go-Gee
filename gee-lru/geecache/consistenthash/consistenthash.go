package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           //采取依赖注入的方式，允许用于替换成自定义的 Hash 函数，也方便测试时替换，默认为 crc32.ChecksumIEEE 算法。
	replicas int            //虚拟节点倍数
	keys     []int          //哈希环
	hashMap  map[int]string //虚拟节点与真实节点的映射表 hashMap，键是虚拟节点的哈希值，值是真实节点的名称。
}

// New 允许自定义虚拟节点倍数和 Hash 函数。
func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 对每一个真实节点 key，对应创建 m.replicas 个虚拟节点，虚拟节点的名称是：strconv.Itoa(i) + key，即通过添加编号的方式区分不同虚拟节点。
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			fmt.Println("hash===", hash, i, key)
			//使用 m.hash() 计算虚拟节点的哈希值，使用 append(m.keys, hash) 添加到环上。
			m.keys = append(m.keys, hash)
			//在 hashMap 中增加虚拟节点和真实节点的映射关系。
			m.hashMap[hash] = key
		}
	}
	// 环上的哈希值排序。
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	fmt.Println("idx ", idx, len(m.keys), m.keys[idx%len(m.keys)], m.hashMap[m.keys[idx%len(m.keys)]])
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

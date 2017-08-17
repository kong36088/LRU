# LRU
Algorithm LRU

LRU算法（least recently used）GO语言实现。

利用一个map以及双向链表，达到搜索和删除双O（1）。

list头代表最新使用过的元素，链表尾部为最旧的元素。当容量满则删除最旧的元素。

## Parts of Merkel Tree.

Merkel Tree has basically three components.

# 1. Leaf Nodes.
# 2. Non-Leaf Nodes.
# 3. Merkel Root.

## Leaf-Nodes:
Nodes at the bottom of the Merkel-Tree are call Leaf nodes.
This is called Transaction-Hash, also Called TX-ID, Transaction-IDs.
(Every Crypto Transaction in a Block!)

## Non-Leaf Nodes.

Transaction Pairs are hashed Together to create Pair of Non-Lead nodes above the
leaf nodes.
These are called Non-Leaf Nodes because unlike Leaf Nodes they don't have transaction IDs.
Two Leaf-Nodes generate One hashed-Leaf-Node.

## Process of Merkel-Root
1. Eight Transactions [1,2,3,4,5,6,7,8]
2. Four Pairs [[hash(1,2]),hash([3,4]),hash([5,6]),hash([7,8])]
3. Repet until Get One hash!
[hash([1,2,3,4]),hash([5,6,7,8])]
4. last Pair will be Merkel-Root!

## A Merkel-Root is included in the Block-Header of Every Block! on the Block-chain!

## Iterator パターン
何かがたくさん集まっているときに、それを順番に指し示していき、全体をスキャンして繰り返すもの

### Iterator
要素を順番にスキャンしていくインタフェースを定める役割。
サンプルでは、Iteratorインタフェースがこの役割を持つ。
### ConcreateIterator
Iterator役が定めたインタフェースを実装する役割。
この役はスキャンするために必要な情報を持つ必要がある。
### Aggregate
Iterator役を作り出すインタフェースを定める役割。

### ConcreateAggregate
Aggregate役が定めたインタフェースを実装する役割。
ConcreteIterator役のインスタンスを作り出す。


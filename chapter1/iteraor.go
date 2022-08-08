package main

import "fmt"

type Book struct {
	name string
}

func (b Book) getName() string {
	return b.name
}

// Iterator
type Iterator interface {
	hasNext() bool
	next() *Book
}

// Concrete Iterator
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bs *BookShelf) newBookShelfIterator(index int) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bs,
		index:     index,
	}
}

func (bsi *BookShelfIterator) hasNext() bool {
	if bsi.index < bsi.bookShelf.getLength() {
		return true
	}
	return false
}

func (bsi *BookShelfIterator) next() *Book {
	book := bsi.bookShelf.getBookAt(bsi.index)
	bsi.index++
	return book
}

//

// Concrete Aggregate
type BookShelf struct {
	books []*Book
	last  int
}

func newBookShelf() *BookShelf {
	return &BookShelf{
		last: 0,
	}
}

func (bs *BookShelf) getBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) append(book *Book) {
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) getLength() int {
	return bs.last
}

//

// Aggregate
func (bs *BookShelf) iterator() Iterator {
	return bs.newBookShelfIterator(0)
}

func main() {
	bookShelf := newBookShelf()
	bookShelf.append(&Book{"AB"})
	bookShelf.append(&Book{"CD"})
	bookShelf.append(&Book{"EF"})
	it := bookShelf.iterator()
	for it.hasNext() {
		book := it.next()
		fmt.Println(book.getName())
	}
}

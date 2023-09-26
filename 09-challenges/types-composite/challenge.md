# Challenge: Working with Composite Types

Let's exercise our composite type muscles in this challenge by creating a library to house some books by various authors. This library should allow us to look up books by author's name and print out details about the book such as its title and author.

You should make use of functions, structs (and their methods), maps, and slices to accomplish this task.

## Acceptance criteria

1. You have 3 types, an `author` with a `name` field, a `book` with a `title` and `author` fields (use the `author` custom type), and a `library`.
2. You've defined an `addBook` method on `library`
3. You've defined a `lookupByAuthorName` method on `library` that accepts a string representing an author's name and that returns a `[]book`
4. You can initialize a library, authors, and books and fill the library
5. You can perform book lookups by author name
6. You can retrieve a book and print its title and it's author's name

# golang-learning

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Short bio and motivation in learning golang.

# Learned lessons

### arrays - 100.0%

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

</details>

<details>
  <summary><code>func SumAllTails(numbersToSum ...[]int) []int</code></summary>

</details>

### adder_test - 100.0%

<details>
  <summary><code>func Add(x, y int) int</code></summary>

</details>

### hellogo - 87.5%

<details>
  <summary><code>func Hello(name string, language string) string</code></summary>

</details>

### iteration - 100.0%

<details>
  <summary><code>func ExampleRepeat(char string, n int) string</code></summary>

</details>

<details>
  <summary><code>func Repeat(character string) string</code></summary>

</details>

### pointers_errors - 85.7%
VARIABLES

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

TYPES

type Bitcoin int

func (b Bitcoin) String() string

type Stringer interface {
	String() string
}

type Wallet struct {
	// Has unexported fields.
}

func (w *Wallet) Balance() Bitcoin

func (w *Wallet) Deposit(amount Bitcoin)

func (w *Wallet) Withdraw(amount Bitcoin) error
### smi - 100.0%
TYPES

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64

func (r Rectangle) Perimeter() float64

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64
### maps - 83.3%
CONSTANTS

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

TYPES

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error

func (d Dictionary) Delete(word string)

func (d Dictionary) Search(word string) (string, error)

func (d Dictionary) Update(word, definition string) error

type DictionaryErr string

func (e DictionaryErr) Error() string
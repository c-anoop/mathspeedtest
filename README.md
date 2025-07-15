# Math Speed Test

A customizable math speed test generator for kids, written in Go. This tool generates random math problems to help children practice their arithmetic skills.

## Features

- **Two modes of operation:**

  - **Speed Mode**: Comprehensive math practice with addition, subtraction, multiplication, and division
  - **Table Mode**: Focused multiplication table practice (1-10 tables)

- **Customizable question count**: Generate any number of questions you need
- **Clean output format**: Problems are displayed in a compact, easy-to-read format
- **Modern Go implementation**: Uses current best practices for random number generation

## Usage

```bash
./speedtest <operation> <count>
```

### Arguments

- `operation`: Either "speed" or "table"
- `count`: Number of questions to generate (positive integer)

### Examples

```bash
# Generate 5 comprehensive speed test questions (20 problems total)
./speedtest speed 5

# Generate 10 multiplication table questions
./speedtest table 10

# Show usage help
./speedtest
```

## Question Types

### Speed Mode

Each question includes 4 different operations:

- **Addition**: 3-digit + 2-3 digit numbers
- **Subtraction**: Ensures positive results
- **Multiplication**: 3-digit × single digit (2-9)
- **Division**: Guarantees whole number results

### Table Mode

- **Multiplication**: Random combinations using numbers 1-10
- Perfect for practicing multiplication tables

## Building and Running

1. **Build the program:**

   ```bash
   go build -o speedtest speedtest.go
   ```

2. **Run the program:**
   ```bash
   ./speedtest speed 5
   ```

## Sample Output

### Speed Mode

```
Math Speed Test - speed mode (2 questions)
========================================
559 + 953 =
527 - 488 =
821 x 4 =
612 ÷ 9 =
149 + 496 =
445 - 207 =
762 x 6 =
798 ÷ 6 =
```

### Table Mode

```
Math Speed Test - table mode (3 questions)
========================================
1 x 3 =
2 x 5 =
4 x 6 =
```

## Requirements

- Go 1.20+ (uses modern random number generation)

## License

This project is open source and available under the terms specified in the LICENSE file.

# Math Speed Test

A customizable math speed test generator for kids, written in Go. This tool generates random math problems to help children practice their arithmetic skills.

## Features

- **Two modes of operation:**

  - **Speed Mode**: Comprehensive math practice with addition, subtraction, multiplication, and division
  - **Table Mode**: Focused multiplication table practice (1-10 tables)

- **Customizable problem count**: Generate 1-50 individual math problems
- **Two-column layout**: Compact display with problems arranged in columns
- **Clean output format**: Problems are displayed in an organized, easy-to-read format
- **Modern Go implementation**: Uses current best practices for random number generation

## Usage

```bash
./speedtest <operation> <count>
```

### Arguments

- `operation`: Either "speed" or "table"
- `count`: Number of individual problems to generate (1-50)

### Examples

```bash
# Generate 10 individual speed test problems
./speedtest speed 10

# Generate 15 multiplication table problems
./speedtest table 15

# Generate maximum allowed problems (50)
./speedtest speed 50

# Show usage help
./speedtest
```

## Problem Types

### Speed Mode

Individual problems cycling through 4 different operations:

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
   ./speedtest speed 10
   ```

## Sample Output

### Speed Mode (Two-Column Layout)

```
Math Speed Test - speed mode (10 problems)
========================================
487 + 690 =                      775 - 398 =
125 x 6 =                        522 ÷ 9 =
804 + 437 =                      336 - 191 =
866 x 9 =                        696 ÷ 4 =
450 + 903 =                      818 - 589 =
```

### Table Mode (Two-Column Layout)

```
Math Speed Test - table mode (8 problems)
========================================
9 x 7 =                          10 x 9 =
5 x 2 =                          7 x 2 =
4 x 3 =                          1 x 7 =
9 x 7 =                          3 x 7 =
```

## Limitations

- Maximum 50 problems per session
- Count must be a positive integer (1-50)

## Requirements

- Go 1.20+ (uses modern random number generation)

## License

This project is open source and available under the terms specified in the LICENSE file.

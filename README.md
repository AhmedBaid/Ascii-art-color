# ASCII Art Color Project

## **Introduction**
This project extends the previous ASCII art generator by adding functionality to colorize the output. The program allows users to specify colors using a specific flag format. Additionally, it remains compatible with other optional flags and arguments for generating ASCII art banners.

---

## **Objectives**

### **Core Requirements:**
1. **Color Output:**
   - The program must allow users to specify text color using the `--color=<color>` flag.
   - The ASCII art will be displayed in the chosen color.

2. **Usage Message:**
   - If the flag is used in an incorrect format, the program must return the following usage message:
     ```
     Usage: go run . [OPTION] [STRING]
     EX: go run . --color=<color> <something to be colored> "something"
     ```

3. **Compatibility:**
   - The program must support other optional arguments (e.g., additional banners or options) as long as they are correctly formatted.
   - It must also function with a single `[STRING]` argument when no additional options are provided.

4. **Error Handling:**
   - Ensure invalid inputs or flag formats are handled gracefully.

---

## **Instructions**

### **Development Guidelines:**
- **Language:** The project must be written in Go.
- **Good Practices:** Follow Go best practices, including modular code design and clear separation of concerns.
- **Unit Testing:** Test files should be created to perform unit testing for the core functions of the program.

---

## **Usage**

### **Examples:**

#### Example 1: Generating ASCII Art with Color Output
```bash
$ go run . --color=red "hello" standard
```
🔴 The output will be displayed in red.

#### Example 2: Using Different Colors
```bash
$ go run . --color=blue 'Hello There!' shadow
```
🔵 The output will be displayed in blue.

### **Error Message for Incorrect Flag Format:**
```bash
$ go run . --color=red "hello"
Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <something to be colored> "something"
```

---

## **Project Requirements**

### **Allowed Packages:**
- Only standard Go packages are allowed.

### **What You Will Learn:**
1. **ANSI Escape Codes:**
   - Implementing terminal text coloring using ANSI escape codes.
2. **Data Manipulation:**
   - Handling and processing strings and data for ASCII art generation.

---

## **Good Practices**
- **Code Organization:** Separate the program logic into modular functions to ensure readability and maintainability.
- **Error Handling:** Handle all potential errors gracefully, such as invalid flag formats or unsupported colors.
- **Testing:** Create unit tests to validate individual components of the program.

---

## **Conclusion**
This project enhances the ASCII art generator by introducing color functionality while maintaining compatibility with other optional features. By adhering to Go best practices and ensuring robust error handling, this project serves as a stepping stone to learning terminal text styling, data manipulation, and modular programming in Go.

# Ascii-art-color

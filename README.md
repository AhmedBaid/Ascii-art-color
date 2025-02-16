# ASCII Art Output Project

## **Introduction**
This project builds upon the previous ASCII art generator by adding functionality to write the output into a file. The program allows users to specify an output file using a specific flag format. Additionally, it remains compatible with other optional flags and arguments for generating ASCII art banners.

---

## **Objectives**

### **Core Requirements:**
1. **Output File:**
   - The program must allow users to specify the output file using the `--output=<fileName.txt>` flag.
   - The file will contain the ASCII art generated from the input string and banner style.

2. **Usage Message:**
   - If the flag is used in an incorrect format, the program must return the following usage message:

     ```
     Usage: go run . [OPTION] [STRING] [BANNER]
     EX: go run . --output=<fileName.txt> something standard
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

#### Example 1: Generating ASCII Art with Output File
```bash
$ go run . --output=banner.txt "hello" standard
$ cat -e banner.txt
 _          _   _     _  $
| |        | | | |   (_) $
| |     ___| |_| |__  _  $
| |    / _ \ __| '_ \| | $
| |___|  __/ |_| | | | | $
|______\___|\__|_| |_|_| $
$
```

#### Example 2: Using Different Banner Styles
```bash
$ go run . --output=banner.txt 'Hello There!' shadow
$ cat -e banner.txt
_|    _|    _|_|_|_|  _|      _|  $
_|    _|  _|        _|_|    _|_|  $
_|_|_|_|  _|  _|_|      _|_|      $
_|    _|  _|    _|      _|        $
_|    _|    _|_|_|      _|        $
$
```

### **Error Message for Incorrect Flag Format:**
```bash
$ go run . --output=banner.txt "hello"
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```

---

## **Project Requirements**

### **Allowed Packages:**
- Only standard Go packages are allowed.

### **What You Will Learn:**
1. **Go File System (fs) API:**
   - Reading and writing files in Go.
2. **Data Manipulation:**
   - Handling and processing strings and data for ASCII art generation.

---

## **Good Practices**
- **Code Organization:** Separate the program logic into modular functions to ensure readability and maintainability.
- **Error Handling:** Handle all potential errors gracefully, such as file I/O errors or invalid flag formats.
- **Testing:** Create unit tests to validate individual components of the program.

---

## **Conclusion**
This project enhances the ASCII art generator by introducing file output functionality while maintaining compatibility with other optional features. By adhering to Go best practices and ensuring robust error handling, this project serves as a stepping stone to learning file handling, data manipulation, and modular programming in Go.
# Ascii-art-color

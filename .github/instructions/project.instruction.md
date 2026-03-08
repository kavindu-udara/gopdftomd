# PDFtoMD - Project Instructions

## Project Overview

**PDFtoMD** is a Go-based PDF to Markdown converter that intelligently extracts content from PDF files while preserving formatting, including headers, bold/italic text, bullet points, and embedded images. It parses PDF content streams and applies intelligent rules to classify text by size and structure.

**Module:** `github.com/kavindu-udara/gopdftomd`  
**Go Version:** 1.25.0+  
**Primary Dependency:** `pdfcpu` - Advanced PDF manipulation library

---

## Architecture Overview

### Core Components

#### 1. **Text Extraction & Parsing** (`reader.go`)
- **TextRun struct**: Represents individual text elements from the PDF with metadata:
  - `Text`: The actual text content
  - `Font`: Font name (used for formatting detection)
  - `Size`: Font size in points
  - `X, Y`: Position coordinates on the page
  - `IsBold`: Boolean flag for bold formatting
  - `IsItalic`: Boolean flag for italic formatting
  - `IsBullet`: Boolean flag for bullet point detection

- **Key Functions**:
  - `parseContentStream(s string)`: Parses raw PDF content streams using regex patterns to extract text operations
  - `detectFormattingFromFont(fontName string)`: Detects bold/italic from font name patterns
  - `detectBulletPoint(text string)`: Identifies bullet points by character prefixes
  - `decodePDFLiteral(s string)`: Decodes PDF escape sequences
  - `cleanText(s string)`: Normalizes text by removing non-printable characters

**Supported Bullet Characters:** `•`, `◦`, `▪`, `■`, `□`, `◆`, `-`, `*`, `+`

#### 2. **Markdown Conversion** (`markdown.go`)
- **MarkdownSizeRules struct**: Configurable thresholds for text classification
  - `H1Min: 20.0` - Heading 1 minimum size
  - `H2Min: 16.0` - Heading 2 minimum size
  - `H3Min: 13.0` - Heading 3 minimum size
  - `BodyMin: 9.0` - Body text minimum size
  - `IgnoreBelow: 7.5` - Skip tiny footer markers
  - `LineMergeTolerance: 2.0` - Y-distance to merge lines
  - `ParagraphGap: 12.0` - Y-distance for paragraph breaks

- **Key Functions**:
  - `classifyBySize(size float64, rules MarkdownSizeRules)`: Classifies text by font size
  - `formatText(text string, isBold, isItalic bool)`: Applies Markdown formatting
    - Bold + Italic: `***text***`
    - Bold only: `**text**`
    - Italic only: `*text*`
  - `runsToMarkdownWithRules(runs []TextRun, rules MarkdownSizeRules)`: Converts text runs to basic Markdown
  - `runsToMarkdownWithRulesAndImages(runs []TextRun, imageRefs []ImageRef, rules MarkdownSizeRules)`: Converts with image embedding

**Output Format:**
- Headings: `# H1`, `## H2`, `### H3`
- Bullet lists: `- Item text`
- Formatting: `**bold**`, `*italic*`, `***bold italic***`
- Paragraphs: Separated by blank lines based on Y-position gaps

#### 3. **Image Extraction & Embedding** (`images.go`)
- **ImageRef struct**: Represents image references in PDFs:
  - `Name`: Image name/identifier
  - `X, Y`: Position coordinates
  - `Markdown`: Generated Markdown image syntax

- **Key Functions**:
  - `parsePageImageRefs(content string)`: Parses image references using PDF transformation matrices (Tm, cm, Do operations)
  - `resolvePageImageRefs(...)`: Maps extracted images to content pages with relative path resolution
  - Matrix operations: `identityMatrix()`, `multiplyMatrix()` for coordinate transformation

**Naming Convention:** Images are extracted as `<pdfname>_<page>_<ImgX>.<ext>`

#### 4. **PDF Control** (`controllers/pdfController.go`)
- `GetPageCount(context *model.Context)`: Returns total page count from PDF context

#### 5. **Parsing Utilities** (`parser.go`)
- `parseLengthFromDict()`: Extracts `/Length` values from PDF dictionaries
- `parseSimpleDict()`: Parses simple key-value pairs from PDF objects
- `isLikelyText()`: Heuristic to identify text vs. binary data
- `truncateString()`: Utility for string truncation

---

## Building & Running

### ⚠️ **Using Makefile (Recommended)**

```bash
# Build the project
make build

# Build and run with default arguments
make run

# Clean build
make clean  # (if available in Makefile)
```

### Manual Build (Without Makefile)

```bash
# Build the binary
go build -o bin/pdf2md .

# Run the binary
./bin/pdf2md -f <input.pdf> -o markdown -out <output_dir>
```

---

## Command-Line Usage

```bash
./bin/pdf2md -f <PDF_FILE> -o <FORMAT> -out <OUTPUT_DIR>
```

### Flags:
- **`-f`** (required): Input PDF file path
- **`-o`** (optional, default: `markdown`): Output format
  - `markdown` - Markdown files with images
  - `text` - Plain text extraction
  - `json` - JSON format
- **`-out`** (optional, default: `output`): Output directory name

### Example:

```bash
make build
./bin/pdf2md -f model.pdf -o markdown -out extracted_content
```

### Output Structure:
```
mds/
├── output_page_1.md
├── output_page_2.md
└── ...

extracted_content/
└── images/
    ├── model_01_Im0.png
    ├── model_01_Im1.jpg
    └── ...
```

---

## Dependencies

```
github.com/pdfcpu/pdfcpu v0.11.1
  ├── github.com/clipperhouse/uax29/v2 v2.2.0
  ├── github.com/hhrutter/lzw v1.0.0
  ├── github.com/hhrutter/tiff v1.0.2
  ├── github.com/pkg/errors v0.9.1
  ├── golang.org/x/crypto v0.43.0
  ├── golang.org/x/image v0.32.0
  ├── golang.org/x/text v0.30.0
  └── gopkg.in/yaml.v2 v2.4.0
```

---

## Key Workflows

### 1. PDF to Markdown Conversion Flow

```
PDF File
    ↓
[pdfcpu] Extract Content Streams & Images
    ↓
Parse Content Stream → TextRun Array
    ├─ detectFormattingFromFont()
    └─ detectBulletPoint()
    ↓
Classify by Size (H1, H2, H3, Body, Ignore)
    ↓
Format Text (Bold, Italic, Bullet Lists)
    ↓
Resolve Image References
    ↓
Output: Markdown Files + Images Directory
```

### 2. Formatting Detection Strategy

**Bold/Italic Detection:**
- Examines font names for patterns: `bold`, `-b`, `italic`, `oblique`, `-i`
- Font names from PDFs: `Arial-Bold`, `Helvetica-Oblique`, etc.

**Bullet Point Detection:**
- Multiple strategies combined:
  1. Text prefix matching (•, -, *, +, etc.)
  2. Left indentation (if enabled in future)
- Renders as Markdown list items: `- text`

**Size-Based Classification:**
- Headers: Font size ≥ 13pt (3 levels)
- Body text: Font size 9–12.9pt
- Small/ignored: Font size < 7.5pt

### 3. Text Run Merging Logic

- **Line Merge**: Runs with Y-position difference ≤ 2pt are merged on same line
- **Paragraph Break**: Runs with Y-position difference ≥ 12pt start new paragraph
- Preserves layout while avoiding excessive whitespace

---

## Configuration & Customization

### Adjusting Markdown Rules

To customize text classification, modify `DefaultMarkdownSizeRules` in `markdown.go`:

```go
var DefaultMarkdownSizeRules = MarkdownSizeRules{
	H1Min:              20.0,  // Increase for larger headings
	H2Min:              16.0,
	H3Min:              13.0,
	BodyMin:            9.0,
	IgnoreBelow:        7.5,   // Skip very small text (footers, watermarks)
	LineMergeTolerance: 2.0,   // Adjust for line spacing
	ParagraphGap:       12.0,  // Increase for larger gaps between paragraphs
}
```

### Extending Bullet Detection

To add more bullet characters, modify the `bulletChars` slice in `detectBulletPoint()`:

```go
bulletChars := []string{"•", "◦", "▪", "■", "□", "◆", "-", "*", "+", "→"}
```

### Font Detection Patterns

To add more font name patterns in `detectFormattingFromFont()`:

```go
isBold := strings.Contains(fontLower, "bold") || 
          strings.Contains(fontLower, "heavy") ||  // Add custom patterns
          strings.HasSuffix(fontLower, "-b")
```

---

## Testing & Validation

### Test PDFs
The project includes test PDFs in the root directory:
- `model.pdf` - Full multi-page document
- `test.pdf` - Simple test document
- `simple.pdf` - Minimal test case

### Testing Steps:

```bash
# Test basic conversion
make build
./bin/pdf2md -f test.pdf -o markdown -out test_output

# Check output
cat mds/output_page_1.md

# Verify images were extracted
ls -la test_output/images/
```

---

## Debugging & Troubleshooting

### Issues & Solutions

| Issue | Cause | Solution |
|-------|-------|----------|
| Text not detected | Font size outside rules | Adjust `BodyMin`, `H1Min` in MarkdownSizeRules |
| Bold/italic not detected | Unrecognized font naming | Add patterns to `detectFormattingFromFont()` |
| Missing bullet points | Custom bullet chars | Add to `bulletChars` array in `detectBulletPoint()` |
| Images not embedded | Path mismatch | Verify image naming follows convention |
| Garbled text | Font encoding issue | Check PDF with external tool first |

### Debug Tips:

1. **Inspect TextRun data**: Add logging in `parseContentStream()` to print parsed runs
2. **Check font names**: Print font names to identify custom font patterns
3. **Validate image refs**: Check if images match expected naming pattern
4. **Test with simpler PDFs**: Use `simple.pdf` before complex documents

---

## Code Quality & Maintenance

### File Responsibilities

| File | Purpose |
|------|---------|
| `main.go` | Entry point, orchestration, page iteration |
| `reader.go` | PDF content stream parsing, formatting detection |
| `markdown.go` | Text classification, Markdown formatting, output generation |
| `images.go` | Image reference parsing, resolution |
| `parser.go` | Generic parsing utilities |
| `controllers/pdfController.go` | PDF metadata extraction |
| `Makefile` | Build automation |

### Style Notes

- **Error Handling**: Uses panic for fatal errors (development-friendly)
- **Regex Patterns**: Pre-compiled for performance
- **Text Normalization**: Uses `unicode.IsPrint` for validation
- **Coordinate System**: PDFs use Y=0 at bottom; handled internally

---

## Future Enhancements

- [ ] Support for multi-level indentation in nested bullet lists
- [ ] Table detection and conversion
- [ ] Link extraction and preservation
- [ ] Color/highlight preservation
- [ ] Configurable font size thresholds via CLI flags
- [ ] JSON output format completion
- [ ] Batch processing mode
- [ ] Performance optimization for large PDFs

---

## Important Notes

1. **Makefile Usage**: Always use `make build` and `make run` for consistency
2. **Output Directory**: Creates `mds/` and `output/` directories automatically
3. **Idempotency**: Re-running overwrites previous output
4. **Permission Requirements**: Needs read access to PDF and write access to destination
5. **PDF Compatibility**: Works best with text-based PDFs; scanned/image PDFs not recommended

---

## Contact & Development

**Module Reference:** `github.com/kavindu-udara/gopdftomd`

For development, clone the repository and use:
```bash
go mod download  # Install dependencies
make build       # Build binary
```

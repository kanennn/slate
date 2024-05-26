package export

import (
	"bytes"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	slateutils "github.com/kanennn/slate/utils"
)

func ExportPdf(html bytes.Buffer, css[]byte) {
	newBuffer := new(bytes.Buffer)
    newBuffer.WriteString("<html lang=en-US>")
    newBuffer.Write(html.Bytes())
	buf := *newBuffer

	target := proto.TargetCreateTarget{URL: ""}
	page, err := rod.New().MustConnect().Page(target)
	slateutils.Check(err)
	htmlcontent := buf.String()

	// width := float64(2)
	// height := float64(2)

	printer := proto.PagePrintToPDF{
		PreferCSSPageSize: true,
		Landscape: false,
		DisplayHeaderFooter: false,
		PrintBackground: false,
		//Scale: 1,
		//PaperWidth: &width,
    	// PaperHeight (optional) Paper height in inches. Defaults to 11 inches.
    	//PaperHeight: &height,

    	// MarginTop (optional) Top margin in inches. Defaults to 1cm (~0.4 inches).
    	//MarginTop: 1

    	// MarginBottom (optional) Bottom margin in inches. Defaults to 1cm (~0.4 inches).		
		//MarginBottom:1,

    	// MarginLeft (optional) Left margin in inches. Defaults to 1cm (~0.4 inches).		
    	//MarginLeft: 1,

    	// MarginRight (optional) Right margin in inches. Defaults to 1cm (~0.4 inches).		
    	//MarginRight: 1,

		// PageRanges (optional) Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are
		// printed in the document order, not in the order specified, and no
		// more than once.
		// Defaults to empty string, which implies the entire document is printed.
		// The page numbers are quietly capped to actual page count of the
		// document, and ranges beyond the end of the document are ignored.
		// If this results in no pages to print, an error is reported.
		// It is an error to specify a range with start greater than end.
    	//PageRanges: "",

		// HeaderTemplate (optional) HTML template for the print header. Should be valid HTML markup with following
		// classes used to inject printing values into them:
		// - `date`: formatted print date
		// - `title`: document title
		// - `url`: document location
		// - `pageNumber`: current page number
		// - `totalPages`: total pages in the document
		//
		// For example, `<span class=title></span>` would generate span containing the title.
    	//HeaderTemplate: "",

    	// FooterTemplate (optional) HTML template for the print footer. Should use the same format as the `headerTemplate`.
    	//FooterTemplate: "",

	}

	page.SetDocumentContent(htmlcontent)
	page.AddStyleTag("", string(css))
	pdf, err := page.PDF(&printer)
	slateutils.Check(err)
	utils.OutputFile("master.pdf", pdf)
}
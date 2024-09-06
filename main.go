package main

// /home/jonathan/MIAB_2S/HT2_MIA_2S_2024/main.go
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello World")
	dotFileName := "/home/jonathan/MIAB_2S/HT2_MIA_2S_2024/DotFile.dot"

	dotContent := `digraph G {
    node [shape=plain, fontname=Arial];
    Disco1 [label=<
        <TABLE BORDER="1" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
            <TR>
                <TD COLSPAN="7">Disco1.dsk</TD>
            </TR>
            <TR>
                <TD>MBR</TD>
                <TD>Libre<br/>25% del disco</TD>
                <TD COLSPAN="3">
                    <TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0">
                        <TR>
                            <TD COLSPAN="5">Extendida</TD>
                        </TR>
                        <TR>
                            <TD>EBR</TD>
                            <TD>Lógica<br/>10% del Disco</TD>
                            <TD>Libre<br/>10% del Disco</TD>
                            <TD>EBR</TD>
                            <TD>Lógica<br/>10% del Disco</TD>
                        </TR>
                    </TABLE>
                </TD>
                <TD>Primaria<br/>20% del disco</TD>
                <TD>Libre<br/>15% del disco</TD>
            </TR>
		</TABLE>
			>];
		}`
	fmt.Println(dotContent)

	// Create a new file and write the content
	file, err := os.Create(dotFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(dotContent)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("dot", "-Tpng", dotFileName, "-o", "/home/jonathan/MIAB_2S/HT2_MIA_2S_2024/DotFile.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

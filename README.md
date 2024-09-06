# HT2_MIA_2S_2024

## Install Graphviz Ubuntu: 
> ```sudo apt-get install graphviz```

## Create a dot File
> ```file, err := os.Create(dotFileName)```

## Write content: 
> ```_, err = file.WriteString(dotContent)```
## Generate an image using grapghviz
> ```	cmd := exec.Command("dot", "-Tpng", dotFileName, "-o", "/home/jonathan/MIAB_2S/HT2_MIA_2S_2024/DotFile.png")```

## Run Code: 
> ``` go run main.go ```

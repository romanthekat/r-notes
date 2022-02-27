# r-notes

This is a collection of tools I use for my notes in markdown, zettelkasten alike, format.  
A tool relies on `[[wiki-link]]` format to find related notes.

## CLI tools
### Build a tool manually
`go build ./cmd/CMD/`  
`go install ./cmd/CMD/`  
`go run ./cmd/CMD/`

### Regenerate-backlinks
Recalculates backlinks for notes, updates files appending the result to the end of a file.   
For example:
```
...note as is...
## Backlinks
- programming languages [[202012051859]]
- always be coding [[202012141632]]
- keeping context [[202106071713]]
- index for 'criteria to select language' [[202111241342]]
- Jevons paradox [[202201161342]]
- Hyrum Law [[202201161344]]
- Chesterton Fence [[202201242307]]
- The Joel Test [[202201251438]]
```

#### Usage
`regenerate-backlinks "path/to/notes/folder"`

#### Install
`go install github.com/romanthekat/r-notes/cmd/regenerate-backlinks@latest`

---

### Sub-graph
Renders subgraph by provided note file.

#### Usage
`sub-graph -notePath="path/to/note.md" -outputPath="path/to/graph.png"`  
`sub-graph -h`
```
Usage of ./sub-graph:
  -depth int
        graph depth to render (default 2)
  -notePath string
        a path to note file
  -outputPath string
        a path to rendered graph file (default "./")
```

#### Install
`go install github.com/romanthekat/r-notes/cmd/sub-graph@latest`

---


### Full-graph
Renders graph of all notes within provided folder.  
Please note, 700+ notes' graph will be somewhat big.

#### Usage
`full-graph "path/to/notes/folder"`

#### Install
`go install github.com/romanthekat/r-notes/cmd/full-graph@latest`

---


### Outliner
Generates an outline for a note with links.  
For example with 3 levels depth (note -> links -> links of links):
```
---
title: index for 'automatic outliner experiment'
date: 2020-12-17 12:24
tags: #index 
---
# 202012171224 index for 'Automatic outliner experiment'
- Automatic outliner experiment [[202012051850]]  
    - Zettelkasten [[202012051855]]  
        - The Archive [[202012061631]]  
        - Org-roam [[202012061643]]  
        - Note taking [[202012061807]]  
        - Knowledge vs information [[202012111603]]  
    - Programming languages [[202012051859]]  
        - Golang [[202012051900]]  
        - Java [[202012051903]]  
        - Python [[202012051904]]  
        - Kotlin [[202012051905]]  
        - Rust [[202012051906]]  
        - Elixir [[202012051907]]  
        - Nim [[202012051908]]  
        - Ruby [[202012051909]]  
        - Criteria to select language [[202012051910]]  
        - Web developer [[202012051919]]  
        - Crystal [[202012051955]]  
        - sql [[202012122037]]  
        - Alternative to java and spring [[202012122049]]  
        - julia [[202012122057]]  
        - Stackoverflow vacancies by language [[202012122102]]  
        - redlang [[202012122104]]  
        - haxe [[202012122106]]  
        - Pony language [[202012122108]]  
    - Golang [[202012051900]]  
        - Compiled languages [[202012051914]]  
        - Static-typed languages [[202012051915]]  
        - Concurrency and parallelism [[202012051916]]  
        - Native binary [[202012051917]]  
        - Web developer [[202012051919]]  
        - Fast compilation [[202012051927]]  
        - Golang manual [[202012061755]]  
        - Miss in golang [[202012061759]]  
        - Frameworks ratings for Golang [[202012122116]]  
    - Org-mode [[202012052002]]  
        - Emacs [[202012052000]]  
    - Auto outliner results [[202012061624]]  
        - The Archive [[202012061631]]  
```

#### Install
`go install github.com/romanthekat/r-notes/cmd/outliner@latest`

#### Usage
`outliner "path/to/note.md"`  
will generate `/path/to/DATE_ZK_ID.md`
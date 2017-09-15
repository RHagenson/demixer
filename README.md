# demiser

A collection of methods to demix Illumina read data, each method in its own directory.
See each demixer's README for details on its method.

A collection of genomic assemblers using different methodologies each in their own directory.
See each assembler's README for details on its method.

## Meta-tasks
Each of these methods has the option to do the following before their specific methods:
1. Removes any barcodes that might have originally IDed the individuals
2. De-duplicates reads (keeping track of how many duplications there were)

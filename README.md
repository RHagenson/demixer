# demixer

A collection of methods to demix Illumina read data -- each method in its own directory.
See each demixer's README for details on its method.

## Meta-tasks
Each of these methods has the option to do the following before their specific methods:
1. Removes any barcodes that might have originally IDed the individuals
2. De-duplicates reads (keeping track of how many duplications there were)

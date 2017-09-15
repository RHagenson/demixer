# demixer/lzdistance

An Illumina demixer that uses Lempel-Ziv (LZ) distance to build out a tree
then splits the tree into k parts to write k Illumina-formated fastq files.

## Limitations
1. Must know k before using
2. Is not biologically driven
3. Can split the reads into wildly unequal parts
4. Cannot tell you what species each output is, only which reads were related on the tree

## Method Overview


### Input
The method takes the mixed Illumina-formated fastq files and a k value.

### Output
The method outputs k Illumina-formated fastq files.

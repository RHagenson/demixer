# demixer/lzgrammar

An Illumina demixer that uses partial assembly of reads, Lempel-Ziv (LZ) grammar,
and k-means clustering to map reads into k fastq files by species.

## Limitations
1. Have to know k before using
2. Likely cannot differentiate very closely-related species
3. Cannot tell you what species each output is, only which reads clustered

## Method Overview
The reads are initially assembled into as large of pieces as possible using
a highly-conservative similarity cutoff (e.g. one mismatch or extend-only) then
uses LZ grammar and k-means clustering to splits the extended reads into k clusters which hopefully correspond to each of k species.

### Input
The method takes the mixed Illumina-formated fastq files and a k value.

### Output
The method outputs k Illumina-formated fastq files.

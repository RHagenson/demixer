# demixer/ami

An Illumina demixer that uses Average Mutational Information (AMI) with species
references (approximately 5 kb each) to split Illumina mixed output
reads into species-specific reads for each species.

## Limitations
1. Must have approximately a 5 kb reference per potential species
2. Uses entropy (H) and mutual information (MI) rather than biological property
3. Reads should be above 100 bp although smaller might be determinable

## Method Overview
An AMI profile is created for each potential species and for each read to
determine which species the read should go to. AMI is a meta-property that
acts as a genetic signature measuring how much information is shared between
the reference and the read(s) -- the more information is shared, the more
likely it is that the read is part of that species.

### Input
The method takes either species reference files or accessions IDs along with
the mixed Illumina-formated fastq files.

### Output
The method outputs Illumina-formated fastq read files named by the first N
characters of the reference header or by accession ID.

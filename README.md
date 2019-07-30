# demixer

A collection of methods to demix Illumina read data -- each method in its own directory.
See each demixer's README for details on its method.

Eventually there is a plan/hope to create a central demixer that feeds data through to each individual demixer which then sends back its confidence in classifying each read to each reference/species.
The central demixer should then take these calls and confidences to make a final call that has far greater confidence due to multi-point verification to reference/species classification.
As such each individual demixer needs an option to spit back confidence scores as opposed to making the call itself and writing appropriate files.
This should be as simple as each demixer optionally returning a tsv/csv formatted list with an ID column and a column for each reference/species with each row being a new read identifiable by ID.
The central demixer would correlate these scores to one another to match columns between methods since Ref1 by one method might be Ref2 in another (e.g. noteably AMI uses species references, but LZ distance uses a k value so AMI returns Species 1-N columns, while LZ distance returns k 1-N with no explict relationships between Species n and k n).
Once columns across methods are matched, the scores across are collapsed into a single confidence score per reference/species (weighting each method according to its predictive nature) and the final call is made given a user-supplied cutoff value (default 80% confidence or 0.80).
The weighting of methods can be optionally given at CLI via `--weights="<method>=<weight>,<method>=<weight>"`

## Meta-tasks that are not a responsibility
The following tasks should not be considered features of these demixers:

1. Removes any barcodes that might have originally IDed the individuals
2. De-duplicates reads (keeping track of how many duplications there were)

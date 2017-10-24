/*
Package lzdistance uses Lempel-Ziv (LZ) distance to compute how similar two
sequences are. LZ distance has been used to differentiate between species in
literature.

Methodology is derived from:
  Otu, et al. 2003 https://doi.org/10.1093/bioinformatics/btg295

In Otu, et al. 2003 there are four distance calculations discussed:
  1. The raw LZ distance equal to the greatest extension of LZ grammar from
     the two sequences. d(S,Q) where S and Q are Sequences
  2. The normalized raw distance d*(S,Q) = d(S,Q) / max{c(S), c(Q)}
  3. The sum distance equal to the sum of LZ grammar size when
     adding S to Q and adding Q to S. ds(S,Q) where S and Q are Sequences
  4. The normalized sum distance ds*(S,Q) = ds(S,Q) / (1/2[c(SQ)+c(QS)])
*/
package lzdistance

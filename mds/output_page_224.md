218 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.17: The steps in transmitting digital information are shown in the upper system, the Funda-
mental Model of Communication. The symbolic-valued signal s ( m ) forms the message, and it is encoded
into a bit sequence b ( n ) . The indices di er because more than one bit/symbol is usually required to
represent the message by a bitstream. Each bit is represented by an analog signal, transmitted through
the (unfriendly) channel, and received by a matched- lter receiver. From the received bitstream
b
b ( n ) the
received symbolic-valued signal b s ( m ) is derived. The lower block diagram shows an equivalent system
wherein the analog portions are combined and modeled by a transition diagram, which shows how each
transmitted bit could be received. For example, transmitting a 0 results in the reception of a 1 with
probability p (an error) or a 0 with probability 1 p (no error).

Communication theory has been formulated best for symbolic-valued signals. Claude Shannon published
in 1948 The Mathematical Theory of Communication , which became the cornerstone of digital communi-
cation. He showed the power of probabilistic models for symbolic-valued signals, which allowed him to
quantify the information present in a signal. In the simplest signal model, each symbol can occur at index
n with a probability Pr [ a ] , k = f 1 ; : : : ; K g . What this model says is that for each signal value a K -sided
coin is ipped (note that the coin need not be fair). For this model to make sense, the probabilities must be
numbers between zero and one and must sum to one.

0 Pr [ a ] 1 (6.48)

X
Pr [ a ] = 1 (6.49)

This coin- ipping model assumes that symbols occur without regard to what preceding or succeeding symbols
were, a false assumption for typed text. Despite this probabilistic model's over-simplicity, the ideas we
develop here also work when more accurate, but still probabilistic, models are used. The key quantity that
characterizes a symbolic-valued signal is the entropy of its alphabet.

H ( A ) =
X
Pr [ a ] log Pr [ a ] (6.50)

Because we use the base-2 logarithm, entropy has units of bits. For this de nition to make sense, we must
take special note of symbols having probability zero of occurring. A zero-probability symbol never occurs;
thus, we de ne 0log 0 = 0 so that such symbols do not a ect the entropy. The maximum value attainable
by an alphabet's entropy occurs when the symbols are equally likely ( Pr [ a ] = Pr [ a ] ). In this case, the
entropy equals log K . The minimum value occurs when only one symbol occurs; it has probability one of
occurring and the rest have probability zero.
This content is available online at http://cnx.org/content/m0070/2.13/ .
http://www-gap.dcs.st-and.ac.uk/~history/Biographies/Shannon.html

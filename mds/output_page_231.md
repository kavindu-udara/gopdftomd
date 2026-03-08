225 p )
(1 p )
(1 p )
(1 p ) (1 p )
(1 p ) (1 p )
: In this example, the transmitter encodes 0 as 000 . The channel creates an error
(changing a 0 into a 1 ) with probability p . The rst column lists all possible received data words
and the second the probability of each data word being received. The last column shows the
results of the majority-vote decoder. When the decoder produces 0 , it successfully corrected the
errors introduced by the channel (if there were any; the top row corresponds to the case in which
no errors occurred). The error probability of the decoders is the sum of the probabilities when
the decoder produces 1 . (Solution on p. 257.)
Demonstrate mathematically that this claim is indeed true. Is 3 p (1 p ) + p p ?

### 6.26 Block Channel Coding

Because of the higher datarate imposed by the channel coder, the probability of bit error occurring in the
digital channel increases relative to the value obtained when no channel coding is used. The bit interval
duration must be reduced by
in comparison to the no-channel-coding situation, which means the energy
per bit E goes down by the same amount. The bit interval must decrease by a factor of three if the
transmitter is to keep up with the data stream, as illustrated in Figure 6.21.

Point of Interest: It is unlikely that the transmitter's power could be increased to compensate.
Such is the sometimes-unfriendly nature of the real world.

Because of this reduction, the error probability p of the digital channel goes up. The question thus becomes
does channel coding really help: Is the e ective error probability lower with channel coding even though the
error probability for each transmitted bit is larger? The answer is no : Using a repetition code for channel
coding cannot ultimately reduce the probability that a data bit is received in error. The ultimate reason is
the repetition code's ine ciency: transmitting one data bit for every three transmitted is too ine cient for
the amount of error correction provided.

Exercise 6.27 (Solution on p. 257.)
Using MATLAB, calculate the probability a bit is received incorrectly with a three-fold repetition
code. Show that when the energy per bit E is reduced by 1 = 3 that this probability is larger than
the no-coding probability of error.

The repetition code (p. 224) represents a special case of what is known as block channel coding . For every
K bits that enter the block channel coder, it inserts an additional N K error-correction bits to produce a
block of N bits for transmission. We use the notation ( N; K ) to represent a given block code's parameters.
In the three-fold repetition code (p. 224), K = 1 and N = 3 . A block code's coding e ciency E equals
the ratio K=N , and quanti es the overhead introduced by channel coding. The rate at which bits must be
This content is available online at http://cnx.org/content/m0094/2.15/ .

223 Morse and Hu man Codes for American-Roman Alphabet. The % column indicates the
average probability (expressed in percent) of the letter occurring in English. The entropy H ( A ) of the
this source is 4.14 bits. The average Morse codeword length is 2.5 symbols. Adding one more symbol for
the letter separator and converting to bits yields an average codeword length of 5.56 bits. The average
Hu man codeword length is 4.35 bits.

We can, to some extent, correct errors made by the receiver with only the error- lled bit stream emerging
from the digital channel available to us. The idea is for the transmitter to send not only the symbol-derived
bits emerging from the source coder but also additional bits derived from the coder's bit stream. These
additional bits, the error correcting bits , help the receiver determine if an error has occurred in the
data bits (the important bits) or in the error-correction bits. Instead of the communication model shown
previously (Figure 6.17), the transmitter inserts a channel coder before analog modulation, and the receiver
the corresponding channel decoder (Figure 6.20). This block diagram shown there forms the Fundamental
Model of Digital Communication .
Shannon's Noisy Channel Coding Theorem says that if the data aren't transmitted too quickly, that error
correction codes exist that can correct all the bit errors introduced by the channel. Unfortunately, Shannon
did not demonstrate an error correcting code that would achieve this remarkable feat; in fact, no one has
This content is available online at http://cnx.org/content/m10782/2.5/ .

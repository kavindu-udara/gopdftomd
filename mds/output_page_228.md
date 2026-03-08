222 CHAPTER 6. INFORMATION COMMUNICATION

### 6.23 Subtleties of Coding

In the Hu man code, the bit sequences that represent individual symbols can have di ering lengths so the
bitstream index m does not increase in lock step with the symbol-valued signal's index n . To capture how
often bits must be transmitted to keep up with the source's production of symbols, we can only compute
averages. If our source code averages ( A ) bits/symbol and symbols are produced at a rate R , the average
bit rate equals ( A ) R , and this quantity determines the bit interval duration T .

Exercise 6.23 (Solution on p. 257.)
Calculate the relation between T and the average bit rate ( A ) R .

A subtlety of source coding is whether we need commas in the bitstream. When we use an unequal number
of bits to represent symbols, how does the receiver determine when symbols begin and end? If you created a
source code that required a separation marker in the bitstream between symbols, it would be very ine cient
since you are essentially requiring an extra symbol in the transmission stream.

note: A good example of this need is the Morse Code: Between each letter, the telegrapher needs
to insert a pause to inform the receiver when letter boundaries occur.

As shown in this example (Example 6.3), no commas are placed in the bitstream, but you can unambiguously
decode the sequence of symbols from the bitstream. Hu man showed that his (maximally e cient) code had
the pre x property: No code for a symbol began another symbol's code. Once you have the pre x property,
the bitstream is partially self-synchronizing: Once the receiver knows where the bitstream starts, we can
assign a unique and correct symbol sequence to the bitstream.

Exercise 6.24 (Solution on p. 257.)
Sketch an argument that pre x coding, whether derived from a Hu man code or not, will provide
unique decoding when an unequal number of bits/symbol are used in the code.

However, having a pre x code does not guarantee total synchronization: After hopping into the middle of a
bitstream, can we always nd the correct symbol boundaries? The self-synchronization issue does mitigate
the use of e cient source coding algorithms.

Exercise 6.25 (Solution on p. 257.)
Show by example that a bitstream produced by a Hu man code is not necessarily self-synchronizing.
Are xed-length codes self synchronizing?

Another issue is bit errors induced by the digital channel; if they occur (and they will), synchronization can
easily be lost even if the receiver started in synch with the source. Despite the small probabilities of error
o ered by good signal set design and the matched lter, an infrequent error can devastate the ability to
translate a bitstream into a symbolic signal. We need ways of reducing reception errors without demanding
that p be smaller.

Example 6.4
The rst electrical communications system the telegraph was digital. When rst deployed in
1844, it communicated text over wireline connections using a binary code the Morse code to
represent individual letters. To send a message from one place to another, telegraph operators
would tap the message using a telegraph key to another operator, who would relay the message
on to the next operator, presumably getting the message closer to its destination. In short, the
telegraph relied on a network not unlike the basics of modern computer networks. To say it
presaged modern communications would be an understatement. It was also far ahead of some
needed technologies, namely the Source Coding Theorem. The Morse code, shown in Figure 6.19,
was not a pre x code. To separate codes for each letter, Morse code required that a space a
pause be inserted between each letter. In information theory, that space counts as another code
letter, which means that the Morse code encoded text with a three-letter source code: dots, dashes
and space. The resulting source code is not within a bit of entropy, and is grossly ine cient (about
25%). Figure 6.19 shows a Hu man code for English text, which as we know is e cient.
This content is available online at http://cnx.org/content/m0093/2.16/ .

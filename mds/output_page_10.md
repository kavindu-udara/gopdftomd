4 CHAPTER 1. INTRODUCTION : The ASCII translation table shows how standard keyboard characters are represented
by integers. In pairs of columns, this table displays rst the so-called 7-bit code (how many
characters in a seven-bit code?), then the character the number represents. The numeric codes
are represented in hexadecimal (base-16) notation. Mnemonic characters correspond to control
characters, some of which may be familiar (like cr for carriage return) and some not ( bel means
a bell ). green and blue is present. Mathematically,
color pictures are multivalued vector-valued signals: s ( x ) = ( r ( x ) ; g ( x ) ; b ( x )) .
Interesting cases abound where the analog signal depends not on a continuous variable, such as time, but
on a discrete variable. For example, temperature readings taken every hour have continuous analog values,
but the signal's independent variable is (essentially) the integers.

1.2.2 Digital Signals

The word digital means discrete-valued and implies the signal depends on the integers rather than a
continuous variable. Digital information includes numbers and symbols (characters typed on the keyboard,
for example). Computers rely on the digital representation of information to manipulate and transform
information. Symbols do not have a numeric value, however each is typically represented by a unique
number but performing arithmetic with these representations makes no sense. The ASCII character code
shown in Table 1.1 has the upper- and lowercase characters, the numbers, punctuation marks, and various
other symbols represented by a seven-bit integer. For example, the ASCII code represents the letter a as
the number 97 , the letter A with 65 .

### 1.3 Structure of Communication Systems

The fundamental model of communications is portrayed in Figure 1.3 (Fundamental model of communi-
cation). In this fundamental model, each message-bearing signal, exempli ed by s ( t ) , is analog and is a
function of time. A system operates on zero, one, or several signals to produce more signals or to simply
absorb them (Figure 1.4). In electrical engineering, we represent a system as a box, receiving input signals
(usually coming from the left) and producing from them new output signals. This graphical representation
is known as a block diagram . We denote input signals by lines having arrows pointing into the box,
output signals by arrows pointing away. As typi ed by the communications model, how information ows,
how it is corrupted and manipulated, and how it is ultimately received is summarized by interconnecting
block diagrams: The outputs of one or more systems serve as the inputs to others.
This content is available online at http://cnx.org/content/m0002/2.17/ .

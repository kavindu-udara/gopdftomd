128 CHAPTER 4. FREQUENCY DOMAIN

(a) Find this lter's transfer function.
(b) Find the magnitude and phase of this transfer function. How would you characterize this circuit?
(c) Let v ( t ) be a square-wave of period T . What is the Fourier series for the output voltage?
(d) Use Matlab to nd the output's waveform for the cases T = 0 : 01 and T = 2 . What value of T delineates
the two kinds of results you found? The software in fourier2.m might be useful.
(e) Instead of the depicted circuit, the square wave is passed through a system that delays its input, which
applies a linear phase shift to the signal's spectrum. Let the delay be
. Use the transfer function
of a delay to compute using Matlab the Fourier series of the output. Show that the square wave is
indeed delayed.

Problem 4.4: Approximating Periodic Signals
Often, we want to approximate a reference signal by a somewhat simpler signal. To assess the quality of
an approximation, the most frequently used error measure is the mean-squared error. For a periodic signal
s ( t ) ,

=
1

Z

( s ( t ) ~ s ( t )) dt

where s ( t ) is the reference signal and ~ s ( t ) its approximation. One convenient way of nding approximations
for periodic signals is to truncate their Fourier series.

~ s ( t ) =
X
c e

The point of this problem is to analyze whether this approach is the best ( i.e. , always minimizes the mean-
squared error).

(a) Find a frequency-domain expression for the approximation error when we use the truncated Fourier
series as the approximation.
(b) Instead of truncating the series, let's generalize the nature of the approximation to including any set of
2 K + 1 terms: We'll always include the c and the negative indexed term corresponding to c . What
selection of terms minimizes the mean-squared error? Find an expression for the mean-squared error
resulting from your choice.
(c) Find the Fourier series for the signal depicted in Figure 4.21. Use Matlab to nd the truncated
approximation and best approximation involving two terms. Plot the mean-squared error as a function
of K for both approximations.

Figure 4.21

Problem 4.5: Long, Hot Days
The daily temperature is a consequence of several e ects, one of them being the sun's heating. If this were
the dominant e ect, then daily temperatures would be proportional to the number of daylight hours. The
map shows that the hottest day of the year varies around the country, but does not occur on the longest day
of the year: June 20 21 (except in the desert Southwest). The plot details the average daily high temperature

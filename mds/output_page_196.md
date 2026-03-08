190 CHAPTER 5. DIGITAL SIGNAL PROCESSING

processing. First of all, he would section the signal into length- N blocks, and compute its N -point DFT. He
then would discard (zero the spectrum) at half of the frequencies, quantize them to b -bits, and send these
over the network. The receiver would assemble the transmitted spectrum and compute the inverse DFT,
thus reconstituting an N -point block.

(a) At what frequencies should the spectrum be zeroed to minimize the error in this lossy compression
scheme?
(b) The nominal way to represent a signal digitally is to use simple b -bit quantization of the time-domain
waveform. How long should a section be in the proposed scheme so that the required number of
bits/sample is smaller than that nominally required?
(c) Assuming that e ective compression can be achieved, would the proposed scheme yield satisfactory
results?

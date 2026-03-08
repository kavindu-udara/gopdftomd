203

centers. Consequently, networks of antennas sprinkle the countryside (each located on the highest hill
possible) to provide long-distance wireless communications: Each antenna receives energy from one antenna
and retransmits to another. This kind of network is known as a relay network .

### 6.6 The Ionosphere and Communications

If we were limited to line-of-sight communications, long distance wireless communication, like ship-to-shore
communication, would be impossible. At the turn of the century, Marconi, the inventor of wireless telegraphy,
boldly tried such long distance communication without any evidence either empirical or theoretical
that it was possible. When the experiment worked, but only at night, physicists scrambled to determine
why (using Maxwell's equations, of course). It was Oliver Heaviside, a mathematical physicist with strong
engineering interests, who hypothesized that an invisible electromagnetic mirror surrounded the earth.
What he meant was that at optical frequencies (and others as it turned out), the mirror was transparent,
but at the frequencies Marconi used, it re ected electromagnetic radiation back to earth. He had predicted
the existence of the ionosphere, a plasma that encompasses the earth at altitudes h between 80 and 180 km
that reacts to solar radiation: It becomes transparent at Marconi's frequencies during the day, but becomes
a mirror at night when solar radiation diminishes. The maximum distance along the earth's surface that

can be reached by a single ionospheric re ection is 2 R cos
R + h
, which ranges between 2,010 and

3,000 km when we substitute minimum and maximum ionospheric altitudes. This distance does not span
the United States or cross the Atlantic; for transatlantic communication, at least two re ections would be

required. The communication delay encountered with a single re ection in this channel is
2
p Rh + h
,

which ranges between 6.8 and 10 ms, again a small time interval.

### 6.7 Communication with Satellites

Global wireless communication relies on satellites. Here, ground stations transmit to orbiting satellites that
amplify the signal and retransmit it back to earth. Satellites will move across the sky unless they are in
geosynchronous orbits , where the time for one revolution about the equator exactly matches the earth's
rotation time of one day. TV satellites would require the homeowner to continually adjust his or her antenna
if the satellite weren't in geosynchronous orbit. Newton's equations applied to orbiting bodies predict that
the time T for one orbit is related to distance from the earth's center R as

R =

r
(6.20)

where G is the gravitational constant and M the earth's mass. Calculations yield R = 42200 km , which
corresponds to an altitude of 35700 km . This altitude greatly exceeds that of the ionosphere, requiring
satellite transmitters to use frequencies that pass through it. Of great importance in satellite communications
is the transmission delay. The time for electromagnetic elds to propagate to a geosynchronous satellite and
return is 0.24 s, a signi cant delay.

Exercise 6.6 (Solution on p. 255.)
In addition to delay, the propagation attenuation encountered in satellite communication far exceeds
what occurs in ionospheric-mirror based communication. Calculate the attenuation incurred by
radiation going to the satellite (one-way loss) with that encountered by Marconi (total going up
and down). Note that the attenuation calculation in the ionospheric case, assuming the ionosphere
acts like a perfect mirror, is not a straightforward application of the propagation loss formula
(p. 201).
This content is available online at http://cnx.org/content/m0539/2.10/ .
This content is available online at http://cnx.org/content/m0540/2.10/ .

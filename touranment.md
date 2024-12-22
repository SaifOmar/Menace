# Tournament
It's a tournament engine

yeah let's make it for my freinds lol

give yousry he flashed message

## design

1. Every friend should have his own stats
    1. what are these stats? (length,weight,strength,stamina,iq)
    2. special traits ? (obese, unhinged, insane, calm, outplayer, etc...)

2. Matchmaking
    1. How should I go about making the matchmaker,
    fair? random? mixed? winner queue ? other ?

3. Match
    1. How does the match go ? (no idea)
    2. How do you decide the winner (based on stats, but no idea on details)

4. Post Match
    1. What should happen after each match

### Structure

1. Player
    1. stats
        1. name
        2. height
        3. weight
        4. strength
        5. stamina
        6. iq
        7. hp
    2. ablities
        1. passive
        **Probably**Gonna**Start**With**Only*Passive
        2. active
        3. normal attacks(hard to balance due to relying on stats)
    3. custom messages (to pass to the logger so that I can show it for style points)

2. Logger
    No idea how should this be done or even code structure,
    but it should be used to show how the fight is going

3. MatchMaker
    1. every player start with elo(just a number to represent rating)
    2. the matchmaker for the first match should just make a random
    allocations to matches until all players have finished their first match
    3. the new elo for every player should be determined based on
    win expectency (if the player was expected to win
    ( based on stats) he gets more points and vice versa )
    4. based on the new elo the matchmaker should then decide the next match,
        the match maker should then make it so that the fights are
    within the lowest diffrential of elo possilbe
    5. and again the matchmaker should then decide
    the new elo for each player based on expectency

    6. the first player to touch a certain elo threshhold should be
    the winner(and the lowest elo loses and ranks are everything in between)

4. Match
    1. the match should run the passives for both players
    2. there should be some randomnes to matches
    (something like hit and miss or something like this)
    so that the matches aren't just the same everytime
    3. player abilities will be initially all have
    the same cooldown and almost the same effects
    4. the player attributes should be taken
    into considreation to asses damage and fatigue and other stuff
    5. a match should have time and the player that
    dies before that time or has lower hp at that time losese
    6. if the match ends call the logger to
    announce the winner or do something call idk

#### Implementation design

v = 0.0.1

1. Players have attributes that are the stats and abilites
2. The passive and ability methods are given the abilites attributes when called
3. the special attributes should do something (no?)

# Nixeus_Discord_Bot
Discord bot that keeps track of your current stock positions

Nixeus records and keeps track of your active positions

To have Nixeus start recording your position, use the following format:

        -in [ticker name] [amount of shares] [date]
        
For example, if you started a position in 10 shares of AAPL on June 7th, 2014, you would enter:

        -in AAPL 10 6/7/14
        
To have Nixeus record the exiting of your position, use the following format:

        -out [ticker name] [amount of shares] [date]
        
For example, if you exited a position of 10 shares of AAPL on May 21st, 2016, you would enter:

       -out AAPL 10 5/21/16
        
For shorthand, if you would like to use todays date for the [date], use the keyword "today"

An example: 

       -in F 500 today
       
Would start recording a position of 500 shares in F on the currrent date.

Note that the "today" keyword can be used for both entering (-in) and exiting (-out) positions.

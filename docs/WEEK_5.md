# Week 5 report
First peer review, of Wenlei Dai's sc2bot [here](https://github.com/rescawen/Wenlei-Dai-sc2bot-tiralabra/issues/1).

Tested the solver by using it to play an online 15-puzzle game [here](http://www.artbylogic.com/puzzles/numSlider/numberShuffle.htm).
Added highlighting next move to the solution rendering to make this easier.

![solve](https://user-images.githubusercontent.com/4224639/52887978-83e76400-3182-11e9-8efd-d434433fcb25.gif)

Found some potential ways to improve heuristics, such as invert distance and
Kenichiro Takahashi's walking distance<sup>[source](https://web.archive.org/web/20141224035932/http://juropollo.xe0.ru/stp_wd_translation_en.htm)</sup>.
Invert distance could be used in addition to the existing manhattan distance
heuristic, as the max value of two admissible heuristics is also an admissible
heuristic. I might try improving the heuristic by adding invert or walking
distance if I have time.

Next steps:
* Read through first peer review to my project and fix mentioned issues
* Go through code to find places without enough tests or comments
* Maybe improve storage efficiency of data structures (e.g. using uint8 instead of int when the values below 255)
* Add invert distance or walking distance as an additional heuristic (if there's time)
* Finish testing and implementation documents
* Write second peer review

Time spent: 6h

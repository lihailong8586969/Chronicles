<?php


class Ai
{
    public function __construct()
    {
        while (true) {
            $question = fgets(STDIN);
            $answer =  preg_replace("/\x{5417}\?+/u", '!', $question);
            echo $answer . PHP_EOL;
        }

    }
}
new Ai();
(ns clj.day01
  (:require [clojure.string :as str]))


(def data (str/split-lines (slurp "resources/day01-input.txt")))


(def numbers {"one" "1"
              "two" "2"
              "three" "3"
              "four" "4"
              "five" "5"
              "six" "6"
              "seven" "7"
              "eight" "8"
              "nine" "9"
              "zero" "0"})

(defn convert [number]
  (or (numbers number) number))


(defn extract-simple-digits [line]
  (let [digits (re-seq #"\d" line)
        first-num (first digits)
        last-num (last digits)
        num (str first-num last-num)]
    (Integer/parseInt num)))

(defn extract-complex-digits [line]
  (let [digits (re-seq #"(zero|one|two|three|four|five|six|seven|eight|nine|\d)" line)
        first-num (convert (first (first digits)))
        last-num (convert (first (last digits)))
        num (str first-num last-num)]
    (Integer/parseInt num)))


(defn p1 [input]
  (->> input
       (map extract-simple-digits)
       (apply +)))

(defn p2 [input]
  (->> input
       (map extract-complex-digits)
       (apply +)))


(comment

  (->> data
       (map extract-complex-digits)
       (apply +))

  (re-seq #"(zero|one|two|three|four|five|six|seven|eigth|nine|\d)" "dhsdone23")

  (convert "1")
  )

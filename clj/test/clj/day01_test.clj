(ns clj.day01-test
  (:require
   [clojure.test :refer [deftest is]]
   [clj.day01 :as day01]))

(deftest convert-test
    (is (= "1" (day01/convert "one"))))
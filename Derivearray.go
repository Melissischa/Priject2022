package devidearray

import (
	"fmt"
	"strings"
)

//задание: разделить массив на не-пустые строки, в которых попарно из первой строки во вторую переходиn по
//одному элементу:
//"1", "2+3+4+5",  "1+2", "3+4+5", "1+2+3", "4+5"

func partList(l []string) (result string) {
	for i := 1; i < len(l); i++ {
		firstPart := strings.Join(l[0:i], " ")
		secondPart := strings.Join(l[i:], " ")
		result += fmt.Sprintf("(%s, %s)", firstPart, secondPart)
	}
	return
}

// по образцу задания фуyкци:
//func PartList(arr []string) string {

func PartList(arr []string) string {
	str := ""                                               //инициируем пуст. строку
	for i := 0 ; i < len(arr) - 1 ; i++ {                   //цикл для перебора по массиву (вперед) до посл. элемента 
	  str += appended(arr, i)                               //с добавлением этих эл-тов после кажд цикла в массив, в конец)
	}
	fmt.Println(str)                                        //выводим получившуюся строку (в терминал)
	return str                                              //фиксируем это знач. как раз ф-и partList
  }
  func appended(stringSet []string, index int) string {          
	slice1 := make([]string, len(stringSet))                        //создаем срез, из строк, длиной. к-я будет в массиве
	copy(slice1, stringSet)                                         //добавляем послед. элемент
	slice1[index] += ","                                             //после эл-та с опред номером (= со всеми) доб ","
	return fmt.Sprintf("(%s)", strings.Join(slice1, " "))            //рез вып ф-и - строка, форматированная 
  }
//
// Алгоритм Фишера – Йетса (Современная версия)
// Алгоритм Саттоло
//
// https://ru.wikipedia.org/wiki/%D0%A2%D0%B0%D1%81%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5_%D0%A4%D0%B8%D1%88%D0%B5%D1%80%D0%B0_%E2%80%94_%D0%99%D0%B5%D1%82%D1%81%D0%B0
//
package main

import(
	"fmt"
	"time"
)

/*
from random import randrange

def sattoloCycle(items):
    i = len(items)
    while i > 1:
        i = i - 1
        j = randrange(i)  # 0 <= j <= i-1
        items[j], items[i] = items[i], items[j]
    return


// arr - массив для перестановки, N - количество элементов в массиве
void shuffle(int* arr, int N)
{
    // инициализация генератора случайных чисел
    srand(time(NULL));
 
    // реализация алгоритма перестановки
    for (int i = N - 1; i >= 1; i--)
    {
        int j = rand() % (i + 1);
 
        int tmp = arr[j];
        arr[j] = arr[i];
        arr[i] = tmp;
    }
}
 
int main()
{
    int arr[] = { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 };
 
    shuffle(arr, 10);
 
    for (int i = 0; i < 10; i++)
        printf("%d ", arr[i]);
 
    printf("\n");
 
    return 0;
}

*/


// arr - массив для перестановки, N - количество элементов в массиве
func shuffleSattolo(arr *[]int) {
	N := len(*arr)

	// shake the generator!
	SRnd64(time.Now().Unix())

	for i := N-1; i > 1; i-- {
		j := RndBetweenU(0, i-1) // 0 <= j <= i-1
		(*arr)[j],(*arr)[i] = (*arr)[i],(*arr)[j]
	}
}

// arr - массив для перестановки, N - количество элементов в массиве
func shuffleFisherYates(arr *[]int) {
	N := len(*arr)

	// shake the generator!
	SRnd64(time.Now().Unix())

	for i := N-1; i >= 1; i-- {
		j := RndBetweenU(0, i+1)
		(*arr)[j],(*arr)[i] = (*arr)[i],(*arr)[j]
	}
}




func main() {
	var arr1 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	var arr2 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }

	fmt.Println("source:              ",arr1)
	shuffleFisherYates(&arr1)
	fmt.Println("Fisher-Yates shuffle:",arr1)
	fmt.Println()
	fmt.Println("source:         ",arr2)
	shuffleSattolo(&arr2)
	fmt.Println("Sattolo shuffle:",arr2)
	fmt.Println()
}



var (
	gRndSeed uint32 = 1 // последнее случайное число
	was int = 0 // была ли вычислена пара чисел
	r float64= 0 // предыдущее число
)
// Начиная с некоторого целого числа x0 =/= 0, задаваемого при помощи фукнции SRnd(),
// при каждом вызове функции Rnd() происходит вычисление нового псевдослучайного 
// числа на основе предыдущего.
func SRnd64(seed int64) {
	SRnd(uint32(seed))
}

func SRnd(seed uint32) {
	if seed == uint32(0) {
		gRndSeed = uint32(1)
	} else {
		gRndSeed = seed
	}
}
// Метод генерации случайных чисел основанный на эффекет переполнения 32-разрядных целых чисел
// возвращает равномерно распределённое случайное число
func RndU() uint32 {
	gRndSeed = gRndSeed * uint32(1664525) + uint32(1013904223)
	return gRndSeed
}
// генерировать челое число из диапазона
// с типами надо подумать...
//TODO: надо потестировать! 
func RndBetweenU(bottom, top int) (result int) {
	// формула генерации случайных чисел по заданному диапазону
	// где bottom - минимальное число из желаемого диапазона
	// top - верхнаяя граница, ширина выборки
	rnd := int(RndU())
	//rndBetween := bottom + rnd % top - bottom
	div := rnd % top
	diff := top - div
	//fmt.Printf("Result: %d  rnd: %d  botton: %d  top: %d  div: %d  diff: %d\n", rndBetween, rnd, bottom, top, div, diff)
	if diff > bottom {
		result = bottom + div
	} else {
		result = div
	}
	return
}

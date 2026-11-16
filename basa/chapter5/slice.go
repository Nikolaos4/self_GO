// Срезы(слайсы) являются ссылочной структоурой данных. Если изменить массив, который является базой для среза, то изменится и сам срез.

package main
import (
	"fmt"
	"sort"
	"reflect"
)

func main() {
	var slice = []int{2, 1, 4, 3};
	fmt.Println(slice);

	str := "Hello, word";
	fmt.Println(str[:5]);

	//емкость:
	arr := [6]int {1, 2, 3, 4, 5, 6};
	slice_from_arr := arr[1:5];
	fmt.Println(len(slice_from_arr));
	fmt.Println(cap(slice_from_arr)); //Базовый массив имеет 6 элементов. Срез начинается с индекса 1 (второй элемент массива). От этого элемента до конца массива остаётся 5 элементов (индексы 1,2,3,4,5). Это и есть ёмкость. Вы можете расширить слайс обратно вправо (но не влево) до этих 5 элементов без создания нового массива.

	var make_arr []int = make([]int, 3);
	make_arr[0] = 1;
	make_arr[1] = 2;
	make_arr[2] = 3;

	make_arr = append(make_arr, 4);

	//удалить 2 элемент
	make_arr = append(make_arr[:1], make_arr[2:]...);
	fmt.Println(make_arr);

	//копирование
	make_arr_2 := make([]int,3);
	copy(make_arr_2, make_arr);
	fmt.Println(make_arr_2);

	//сортировка
	//sort.Ints(), sort.Float64s, sort.Strings
	slice_of_number := []int {3, 1, 8, 2, 6};
	sort.Ints(slice_of_number);
	fmt.Println(slice_of_number);

	//поиск
	//sort.SearchInts(), sort.SearchFloat64s(), sort.SearchStrings
	fmt.Println(sort.SearchInts(slice_of_number, 1));

	//сортировка в обратном порядке
	//sort.IntSlice/sort.Float64Slice/sort.StringSlice
	sort.Sort(sort.Reverse(sort.IntSlice(slice_of_number)))
	fmt.Println(slice_of_number);

	//сравнение 
	fmt.Println(reflect.DeepEqual(make_arr_2, slice_of_number));
}
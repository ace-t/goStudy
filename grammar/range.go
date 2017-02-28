package main
import "fmt"
func range_test() {
// 다음 예시는 슬라이스에 있는 숫자들을 더하기 위해 range를 사용합니다. 배열에서도 똑같이 동작합니다.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
// range를 배열과 슬라이스에서 사용하면 각 원소의 인덱스와 값을 반환합니다.
// 위 예시에선 인덱스가 필요 없었기 때문에 공백 식별자 _로 인덱스를 무시했습니다. 


// 다음 예시처럼 인덱스가 필요할 때도 있습니다.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
// 맵에서 range는 key/value 쌍들을 순회합니다.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

// range로 map의 키값들만 순회할 수도 있습니다.
	for k := range kvs {
		fmt.Println("key:", k)
	}

// 문자열에서 range는 유니코드 코드 포인트들을 순회합니다. 첫번째 값은 rune의 시작 바이트 위치이며 두번째 값은 rune값 자체입니다.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

/*
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
*/
}

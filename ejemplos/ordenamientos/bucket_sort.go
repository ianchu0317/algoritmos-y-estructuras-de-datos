package ordenamientos

func BucketSort(array []int) []int {
	// Pre condicion: El arreglo recibido tiene elementos x en [0,100]
	maxNum := 100

	// 1- Creo mis baldes (10 baldes) O(b)
	numBuckets := 10
	buckets := make([][]int, numBuckets)

	// 2- Insertar elementos en buckets O(n*b)
	for _, num := range array {
		indxBucket := getBucketNumber(num, numBuckets, maxNum)
		buckets[indxBucket] = append(buckets[indxBucket], num)
	}

	//fmt.Println("Buckets son:", buckets)

	// 3- Ordenar elementos n log (n/b)
	for i := range buckets {
		buckets[i] = MergeSort(buckets[i])
	}

	// 4- Insertar ordenadamente O(n)
	i := 0
	for _, bucket := range buckets {
		for _, num := range bucket {
			array[i] = num
			i++
		}
	}

	return array
}

func getBucketNumber(num, numBuckets, maxNum int) int {
	bucketSize := maxNum / numBuckets

	if bucketSize == 0 {
		return 0
	}

	return num / bucketSize
}

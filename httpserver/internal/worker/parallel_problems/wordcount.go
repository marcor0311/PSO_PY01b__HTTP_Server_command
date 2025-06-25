package parallelproblems

import (
    "sync"
    "io"
    "bufio"
    "strings"
)

func WordCount(r io.Reader) (int, error) {
    data, err := io.ReadAll(r)
    if err != nil {
        return 0, err
    }

    text := string(data)
    lines := strings.Split(text, "\n")

    var wg sync.WaitGroup
    wordCounts := make(chan int, len(lines))

    for _, line := range lines {
        wg.Add(1)
        go func(l string) {
            defer wg.Done()
            scanner := bufio.NewScanner(strings.NewReader(l))
            scanner.Split(bufio.ScanWords)

            count := 0
            for scanner.Scan() {
                word := scanner.Text()
                if len(strings.TrimSpace(word)) > 0 {
                    count++
                }
            }
            wordCounts <- count
        }(line)
    }

    wg.Wait()
    close(wordCounts)

    total := 0
    for c := range wordCounts {
        total += c
    }

    return total, nil
}


class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    canPermutePalindrome(s: string): boolean {
        const fm: Record<string, number> = {}

        for (let i = 0; i < s.length; i++) {
            if (fm[s[i]] == undefined) {
                fm[s[i]] = 1
                continue
            }

            fm[s[i]]++
        }

        let evenCount = 0
        let oddCount = 0

        for (let key in fm) {
            if (fm[key] % 2 == 0) {
                oddCount++
                continue
            }

            evenCount++
        }
       
       return evenCount <= 1
    }
}

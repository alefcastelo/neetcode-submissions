class Solution {
    /**
     * @param {string[]} sentence1
     * @param {string[]} sentence2
     * @param {string[][]} similarPairs
     * @return {boolean}
     */
    areSentencesSimilar(
        sentence1: string[],
        sentence2: string[],
        similarPairs: string[][],
    ) {

        if (sentence1.length != sentence2.length) {
            return false
        }

        const pairsMap: Record<string, Set<string>> = {}

        for(const pairs of similarPairs) {
            if (pairsMap[pairs[0]] === undefined) {
                pairsMap[pairs[0]] = new Set(pairs)
            } else {
                pairs.map(val => {
                    pairsMap[pairs[0]].add(val)   
                })
            }

            if (pairsMap[pairs[1]] === undefined) {
                pairsMap[pairs[1]] = new Set(pairs)
            } else {
                pairs.map(val => {
                    pairsMap[pairs[1]].add(val)   
                })
            }

        }
            console.log(pairsMap)

        for (let i = 0; i < similarPairs.length; i++) {
            const w1 = sentence1[i]
            const w2 = sentence2[i]

            if (w1 == w2) {
                continue
            }


            if ((pairsMap[w1] && !pairsMap[w1].has(w2)) || 
                (pairsMap[w2] && !pairsMap[w2].has(w1))) {
                return false;
            }
            
        }

        return true;
    }
}



namespace LeetCode;

public class TopInterview150
{
    // https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
    // Given an array nums of size n, return the majority element.
    //
    // The majority element is the element that appears more than ⌊n / 2⌋ times.
    // You may assume that the majority element always exists in the array.
    public int MajorityElement(int[] nums)
    {
        var m = new Dictionary<int, int>();
        foreach (var num in nums)
        {
            if (!m.TryAdd(num, 1))
            {
                m[num]++;
            }
        }

        var majority= nums[0];
        foreach (var (k, v) in m)
        {
            if (v > m[majority])
            {
                majority = k;
            }
        }
        
        return majority;
    }
    
    // https://leetcode.com/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150
    // Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
    // Symbol       Value
    // ------------------
    // I             1
    // V             5
    // X             10
    // L             50
    // C             100
    // D             500
    // M             1000
    public int RomanToInt(string s)
    {
        var sum = 0;
        var lastNum = 1;
        for (var i = s.Length - 1; i >= 0; i--)
        {
            // IV
            var currentNum = s[i] switch
            {
                'I' => 1,
                'V' => 5,
                'X' => 10,
                'L' => 50,
                'C' => 100,
                'D' => 500,
                'M' => 1000,
                _ => 0
            };
            if (currentNum >= lastNum)
            {
                sum += currentNum;
            }
            else
            {
                sum -= currentNum;
            }
            
            lastNum = currentNum;
        }
        
        return sum;
    }
    
    // https://leetcode.com/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150
    public int LengthOfLastWord(string s)
    {
        // "a" - ok
        // "ab" - ok
        // "a " - ok
        // " a" - ok
        // "a bc " - ok
        
        var i = s.Length - 1; // 1
        var end = 0; // 3
        var start = 0; // 0
        while (i >= 0)
        {
            if (s[i] == ' ')
            {
                if (end != 0)
                {
                    return end - i;
                }
            }
            else
            {
                if (end == 0)
                {
                    end = i;
                }
            }

            i--;
        }

        return end + 1;
    }
    
    // https://leetcode.com/problems/is-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
    public bool IsSubsequence(string s, string t) {
        // s = "", t = ""
        // s = "a", t = "bcd"
        // s = "a", t = "bad"
        // s = "cat", t = "cactus"
        // s = "cat", t = "atc"
        
        if (s.Length == 0)
        {
            return true;
        }

        var i = 0;
        foreach (var c in t)
        {
            if (s[i] != c) continue;
            
            i++;
            if (i == s.Length)
            {
                return true;
            }
        }
        
        return false;
    }
    
    
    // https://leetcode.com/problems/ransom-note/description/?envType=study-plan-v2&envId=top-interview-150
    public bool CanConstruct(string ransomNote, string magazine)
    {
        const int start = 'a';
        const int end = 'z';
        var m = new int[end-start+1];
        foreach (var c in magazine)
        {
            m[c - start]++;
        }

        foreach (var c in ransomNote)
        {
            var i = c - start;
            if (m[i] > 0)
            {
                m[i]--;
            }
            else
            {
                return false;
            }
        }

        return true;
    }
    
    // https://leetcode.com/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=top-interview-150
    public bool IsIsomorphic(string s, string t)
    {
        // add   , egg   => true
        // f11   , b23   => false
        // paper , title => true
        // badc  , baba => true
        var st = new Dictionary<char, char>();

        for (var i = 0; i < s.Length; i++)
        {
            if (st.ContainsKey(s[i]))
            {
                if (st[s[i]] == t[i])
                {
                    continue;
                }
                
                return false;
            }
            
            st.Add(s[i], t[i]);
        }
        
        st.Clear();
        
        for (var i = 0; i < t.Length; i++)
        {
            if (st.ContainsKey(t[i]))
            {
                if (st[t[i]] == s[i])
                {
                    continue;
                }
                
                return false;
            }
            
            st.Add(t[i], s[i]);
        }

        return true;
    }
}
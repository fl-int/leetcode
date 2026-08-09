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
}
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
}
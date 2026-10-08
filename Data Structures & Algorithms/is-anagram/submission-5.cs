public class Solution {
    public bool IsAnagram(string s, string t) {
        if (s.Length != t.Length) return false;
        if (s.Equals(t)) return true;

        int[] frequency = new int[26];
        for (int i = 0; i < s.Length; i++)
        {
            frequency[s[i] - 'a']++;
            frequency[t[i] - 'a']--;
        }

        return frequency.All(j => j == 0);
    }
}

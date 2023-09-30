// generated by Textmapper; DO NOT EDIT

#include "json_lexer.h"

#include "absl/log/log.h"
#include "absl/strings/match.h"

namespace json {
namespace {
constexpr int tmNumClasses = 27;

// Latin-1 characters.
constexpr uint8_t tmRuneClass[] = {
    1,  1,  1,  1,  1,  1,  1,  1,  1,  2,  2,  1,  1,  2,  1,  1,  1,  1,  1,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  2,  3,  4,  1,  1,  1,
    1,  1,  1,  1,  5,  6,  7,  8,  9,  10, 11, 12, 12, 12, 12, 12, 12, 12, 12,
    12, 13, 1,  1,  1,  1,  1,  1,  14, 15, 15, 15, 16, 15, 17, 17, 17, 17, 17,
    17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 18, 19, 20, 1,
    1,  1,  15, 21, 15, 15, 16, 21, 17, 17, 17, 17, 17, 17, 17, 22, 17, 17, 17,
    22, 17, 22, 23, 17, 17, 17, 17, 17, 24, 1,  25, 1,  1,  1,  1,  1,  1,  1,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  26,
    1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  26, 1,  1,  1,  1,  26, 1,  1,  1,
    1,  1,  26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 1,  26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
    1,  26, 26, 26, 26, 26, 26, 26, 26,
};

constexpr int tmRuneClassLen = 256;
constexpr int tmFirstRule = -4;

struct MapRange {
  uint32_t lo;
  uint32_t hi;
  uint8_t default_val;
  std::vector<uint8_t> val;
};

const std::vector<MapRange> tmRuneRanges = {
    {256, 706, 26, {}},
    {710, 722, 26, {}},
    {736, 741, 26, {}},
    {748,
     751,
     26,
     {
         26,
         1,
     }},
    {880,
     930,
     26,
     {
         26, 26, 26, 26, 26, 1, 26, 26, 1, 1,  26, 26, 26, 26, 1,
         26, 1,  1,  1,  1,  1, 1,  26, 1, 26, 26, 26, 1,  26, 1,
     }},
    {931, 1014, 26, {}},
    {1015, 1154, 26, {}},
    {1162, 1328, 26, {}},
    {1329, 1367, 26, {}},
    {1369,
     1417,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
         1,
     }},
    {1488, 1515, 26, {}},
    {1519, 1523, 26, {}},
    {1568, 1611, 26, {}},
    {1646,
     1748,
     26,
     {
         26,
         26,
         1,
     }},
    {1749, 1750, 26, {}},
    {1765, 1767, 26, {}},
    {1774, 1776, 26, {}},
    {1786,
     1792,
     26,
     {
         26,
         26,
         26,
         1,
         1,
     }},
    {1808,
     1840,
     26,
     {
         26,
         1,
     }},
    {1869, 1958, 26, {}},
    {1969, 1970, 26, {}},
    {1994, 2027, 26, {}},
    {2036,
     2070,
     26,
     {
         26,
         26,
         1,
         1,
         1,
         1,
         26,
         1,
         1,
         1,
         1,
         1,
     }},
    {2074, 2075, 26, {}},
    {2084,
     2089,
     26,
     {
         26,
         1,
         1,
         1,
     }},
    {2112, 2137, 26, {}},
    {2144, 2155, 26, {}},
    {2160, 2184, 26, {}},
    {2185, 2191, 26, {}},
    {2208, 2250, 26, {}},
    {2308, 2362, 26, {}},
    {2365, 2366, 26, {}},
    {2384,
     2402,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
         1,
         1,
     }},
    {2417, 2433, 26, {}},
    {2437, 2445, 26, {}},
    {2447,
     2473,
     26,
     {
         26,
         26,
         1,
         1,
     }},
    {2474,
     2494,
     26,
     {
         26, 26, 26, 26, 26, 26, 26, 1, 26, 1, 1, 1, 26, 26, 26, 26, 1, 1, 1,
     }},
    {2510, 2511, 26, {}},
    {2524,
     2530,
     26,
     {
         26,
         26,
         1,
     }},
    {2544, 2546, 26, {}},
    {2556, 2557, 26, {}},
    {2565, 2571, 26, {}},
    {2575,
     2601,
     26,
     {
         26,
         26,
         1,
         1,
     }},
    {2602,
     2618,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {2649,
     2655,
     26,
     {
         26,
         26,
         26,
         26,
         1,
     }},
    {2674, 2677, 26, {}},
    {2693, 2702, 26, {}},
    {2703,
     2729,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {2730,
     2750,
     26,
     {
         26, 26, 26, 26, 26, 26, 26, 1, 26, 26, 1, 26, 26, 26, 26, 26, 1, 1, 1,
     }},
    {2768, 2769, 26, {}},
    {2784, 2786, 26, {}},
    {2809, 2810, 26, {}},
    {2821, 2829, 26, {}},
    {2831,
     2857,
     26,
     {
         26,
         26,
         1,
         1,
     }},
    {2858,
     2878,
     26,
     {
         26, 26, 26, 26, 26, 26, 26, 1, 26, 26, 1, 26, 26, 26, 26, 26, 1, 1, 1,
     }},
    {2908,
     2914,
     26,
     {
         26,
         26,
         1,
     }},
    {2929, 2930, 26, {}},
    {2947,
     2955,
     26,
     {
         26,
         1,
     }},
    {2958,
     3002,
     26,
     {
         26, 26, 26, 1, 26, 26, 26, 26, 1, 1, 1,  26, 26, 1, 26, 1,
         26, 26, 1,  1, 1,  26, 26, 1,  1, 1, 26, 26, 26, 1, 1,  1,
     }},
    {3024, 3025, 26, {}},
    {3077, 3085, 26, {}},
    {3086,
     3113,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {3114, 3130, 26, {}},
    {3133, 3134, 26, {}},
    {3160,
     3170,
     26,
     {
         26,
         26,
         26,
         1,
         1,
         26,
         1,
         1,
     }},
    {3200,
     3213,
     26,
     {
         26,
         1,
         1,
         1,
         1,
     }},
    {3214,
     3241,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {3242, 3252, 26, {}},
    {3253,
     3262,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
         1,
         1,
     }},
    {3293,
     3298,
     26,
     {
         26,
         26,
         1,
     }},
    {3313, 3315, 26, {}},
    {3332, 3341, 26, {}},
    {3342,
     3387,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {3389, 3390, 26, {}},
    {3406,
     3415,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
     }},
    {3423, 3426, 26, {}},
    {3450, 3456, 26, {}},
    {3461, 3479, 26, {}},
    {3482, 3506, 26, {}},
    {3507, 3516, 26, {}},
    {3517,
     3527,
     26,
     {
         26,
         1,
         1,
     }},
    {3585, 3633, 26, {}},
    {3634, 3636, 26, {}},
    {3648, 3655, 26, {}},
    {3713,
     3748,
     26,
     {
         26,
         26,
         1,
         26,
         1,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {3749,
     3761,
     26,
     {
         26,
         1,
     }},
    {3762, 3764, 26, {}},
    {3773,
     3783,
     26,
     {
         26,
         1,
         1,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {3804, 3808, 26, {}},
    {3840, 3841, 26, {}},
    {3904, 3912, 26, {}},
    {3913, 3949, 26, {}},
    {3976, 3981, 26, {}},
    {4096, 4139, 26, {}},
    {4159, 4160, 26, {}},
    {4176, 4182, 26, {}},
    {4186,
     4199,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
         1,
         26,
         1,
         1,
         1,
     }},
    {4206,
     4226,
     26,
     {
         26,
         26,
         26,
         1,
         1,
         1,
         1,
     }},
    {4238, 4239, 26, {}},
    {4256, 4294, 26, {}},
    {4295,
     4347,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
         26,
         1,
         1,
     }},
    {4348, 4681, 26, {}},
    {4682,
     4745,
     26,
     {
         26, 26, 26, 26, 1, 1,  26, 26, 26, 26, 26,
         26, 26, 1,  26, 1, 26, 26, 26, 26, 1,  1,
     }},
    {4746,
     4785,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {4786,
     4823,
     26,
     {
         26, 26, 26, 26, 1, 1,  26, 26, 26, 26, 26,
         26, 26, 1,  26, 1, 26, 26, 26, 26, 1,  1,
     }},
    {4824, 4881, 26, {}},
    {4882,
     4955,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {4992, 5008, 26, {}},
    {5024, 5110, 26, {}},
    {5112, 5118, 26, {}},
    {5121, 5741, 26, {}},
    {5743, 5760, 26, {}},
    {5761, 5787, 26, {}},
    {5792, 5867, 26, {}},
    {5873, 5881, 26, {}},
    {5888, 5906, 26, {}},
    {5919, 5938, 26, {}},
    {5952, 5970, 26, {}},
    {5984, 5997, 26, {}},
    {5998, 6001, 26, {}},
    {6016, 6068, 26, {}},
    {6103,
     6109,
     26,
     {
         26,
         1,
         1,
         1,
         1,
     }},
    {6176, 6265, 26, {}},
    {6272,
     6313,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {6314,
     6390,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
     }},
    {6400, 6431, 26, {}},
    {6480, 6510, 26, {}},
    {6512, 6517, 26, {}},
    {6528, 6572, 26, {}},
    {6576, 6602, 26, {}},
    {6656, 6679, 26, {}},
    {6688, 6741, 26, {}},
    {6823, 6824, 26, {}},
    {6917, 6964, 26, {}},
    {6981, 6989, 26, {}},
    {7043, 7073, 26, {}},
    {7086, 7088, 26, {}},
    {7098, 7142, 26, {}},
    {7168, 7204, 26, {}},
    {7245, 7248, 26, {}},
    {7258, 7294, 26, {}},
    {7296, 7305, 26, {}},
    {7312, 7355, 26, {}},
    {7357, 7360, 26, {}},
    {7401,
     7616,
     26,
     {
         26, 26, 26, 26, 1, 26, 26, 26, 26, 26, 26, 1,
         26, 26, 1,  1,  1, 26, 1,  1,  1,  1,  1,
     }},
    {7680, 7958, 26, {}},
    {7960,
     8006,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {8008,
     8024,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {8025,
     8062,
     26,
     {
         26,
         1,
         26,
         1,
         26,
         1,
     }},
    {8064, 8117, 26, {}},
    {8118,
     8141,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         1,
         1,
         1,
         26,
         26,
         26,
         1,
     }},
    {8144,
     8156,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {8160, 8173, 26, {}},
    {8178,
     8189,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {8305, 8306, 26, {}},
    {8319, 8320, 26, {}},
    {8336, 8349, 26, {}},
    {8450,
     8468,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         26,
         1,
         1,
     }},
    {8469,
     8478,
     26,
     {
         26,
         1,
         1,
         1,
     }},
    {8484,
     8506,
     26,
     {
         26,
         1,
         26,
         1,
         26,
         1,
         26,
         26,
         26,
         26,
         1,
     }},
    {8508, 8512, 26, {}},
    {8517, 8522, 26, {}},
    {8526, 8527, 26, {}},
    {8579, 8581, 26, {}},
    {11264, 11493, 26, {}},
    {11499,
     11508,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
         1,
     }},
    {11520, 11558, 26, {}},
    {11559,
     11624,
     26,
     {
         26,
         1,
         1,
         1,
         1,
         1,
         26,
         1,
         1,
     }},
    {11631, 11632, 26, {}},
    {11648, 11671, 26, {}},
    {11680,
     11743,
     26,
     {
         26, 26, 26, 26, 26, 26, 26, 1,  26, 26, 26, 26, 26, 26,
         26, 1,  26, 26, 26, 26, 26, 26, 26, 1,  26, 26, 26, 26,
         26, 26, 26, 1,  26, 26, 26, 26, 26, 26, 26, 1,  26, 26,
         26, 26, 26, 26, 26, 1,  26, 26, 26, 26, 26, 26, 26, 1,
     }},
    {11823, 11824, 26, {}},
    {12293, 12295, 26, {}},
    {12337, 12342, 26, {}},
    {12347,
     12439,
     26,
     {
         26,
         26,
         1,
         1,
         1,
         1,
     }},
    {12445,
     12539,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {12540, 12544, 26, {}},
    {12549, 12592, 26, {}},
    {12593, 12687, 26, {}},
    {12704, 12736, 26, {}},
    {12784, 12800, 26, {}},
    {13312, 19904, 26, {}},
    {19968, 42125, 26, {}},
    {42192, 42238, 26, {}},
    {42240, 42509, 26, {}},
    {42512, 42528, 26, {}},
    {42538, 42540, 26, {}},
    {42560, 42607, 26, {}},
    {42623, 42654, 26, {}},
    {42656, 42726, 26, {}},
    {42775, 42784, 26, {}},
    {42786, 42889, 26, {}},
    {42891, 42955, 26, {}},
    {42960,
     42970,
     26,
     {
         26,
         26,
         1,
         26,
         1,
     }},
    {42994, 43010, 26, {}},
    {43011,
     43043,
     26,
     {
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         26,
         1,
     }},
    {43072, 43124, 26, {}},
    {43138, 43188, 26, {}},
    {43250, 43256, 26, {}},
    {43259,
     43263,
     26,
     {
         26,
         1,
     }},
    {43274, 43302, 26, {}},
    {43312, 43335, 26, {}},
    {43360, 43389, 26, {}},
    {43396, 43443, 26, {}},
    {43471, 43472, 26, {}},
    {43488,
     43504,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {43514,
     43561,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {43584,
     43596,
     26,
     {
         26,
         26,
         26,
         1,
     }},
    {43616, 43639, 26, {}},
    {43642,
     43696,
     26,
     {
         26,
         1,
         1,
         1,
     }},
    {43697,
     43715,
     26,
     {
         26,
         1,
         1,
         1,
         26,
         26,
         1,
         1,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
         26,
         1,
     }},
    {43739,
     43755,
     26,
     {
         26,
         26,
         26,
         1,
         1,
     }},
    {43762, 43765, 26, {}},
    {43777,
     43799,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {43808,
     43867,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {43868, 43882, 26, {}},
    {43888, 44003, 26, {}},
    {44032, 55204, 26, {}},
    {55216, 55239, 26, {}},
    {55243, 55292, 26, {}},
    {63744, 64110, 26, {}},
    {64112, 64218, 26, {}},
    {64256, 64263, 26, {}},
    {64275, 64280, 26, {}},
    {64285,
     64297,
     26,
     {
         26,
         1,
     }},
    {64298, 64311, 26, {}},
    {64312,
     64434,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         1,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {64467, 64830, 26, {}},
    {64848, 64912, 26, {}},
    {64914, 64968, 26, {}},
    {65008, 65020, 26, {}},
    {65136,
     65277,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {65313, 65339, 26, {}},
    {65345, 65371, 26, {}},
    {65382, 65471, 26, {}},
    {65474,
     65501,
     26,
     {
         26, 26, 26, 26, 26, 26, 1,  1,  26, 26, 26, 26,
         26, 26, 1,  1,  26, 26, 26, 26, 26, 26, 1,  1,
     }},
    {65536, 65548, 26, {}},
    {65549, 65575, 26, {}},
    {65576, 65595, 26, {}},
    {65596,
     65614,
     26,
     {
         26,
         26,
         1,
     }},
    {65616, 65630, 26, {}},
    {65664, 65787, 26, {}},
    {66176, 66205, 26, {}},
    {66208, 66257, 26, {}},
    {66304, 66336, 26, {}},
    {66349, 66369, 26, {}},
    {66370, 66378, 26, {}},
    {66384, 66422, 26, {}},
    {66432, 66462, 26, {}},
    {66464, 66500, 26, {}},
    {66504, 66512, 26, {}},
    {66560, 66718, 26, {}},
    {66736, 66772, 26, {}},
    {66776, 66812, 26, {}},
    {66816, 66856, 26, {}},
    {66864, 66916, 26, {}},
    {66928, 66939, 26, {}},
    {66940, 66955, 26, {}},
    {66956,
     66978,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {66979, 66994, 26, {}},
    {66995,
     67005,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {67072, 67383, 26, {}},
    {67392, 67414, 26, {}},
    {67424, 67432, 26, {}},
    {67456,
     67505,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {67506, 67515, 26, {}},
    {67584,
     67638,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         1,
         26,
         1,
     }},
    {67639,
     67670,
     26,
     {
         26,
         26,
         1,
         1,
         1,
         26,
         1,
         1,
     }},
    {67680, 67703, 26, {}},
    {67712, 67743, 26, {}},
    {67808, 67827, 26, {}},
    {67828, 67830, 26, {}},
    {67840, 67862, 26, {}},
    {67872, 67898, 26, {}},
    {67968, 68024, 26, {}},
    {68030, 68032, 26, {}},
    {68096, 68097, 26, {}},
    {68112,
     68150,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         1,
     }},
    {68192, 68221, 26, {}},
    {68224, 68253, 26, {}},
    {68288, 68296, 26, {}},
    {68297, 68325, 26, {}},
    {68352, 68406, 26, {}},
    {68416, 68438, 26, {}},
    {68448, 68467, 26, {}},
    {68480, 68498, 26, {}},
    {68608, 68681, 26, {}},
    {68736, 68787, 26, {}},
    {68800, 68851, 26, {}},
    {68864, 68900, 26, {}},
    {69248, 69290, 26, {}},
    {69296, 69298, 26, {}},
    {69376, 69405, 26, {}},
    {69415, 69416, 26, {}},
    {69424, 69446, 26, {}},
    {69488, 69506, 26, {}},
    {69552, 69573, 26, {}},
    {69600, 69623, 26, {}},
    {69635, 69688, 26, {}},
    {69745,
     69750,
     26,
     {
         26,
         26,
         1,
         1,
     }},
    {69763, 69808, 26, {}},
    {69840, 69865, 26, {}},
    {69891, 69927, 26, {}},
    {69956,
     69960,
     26,
     {
         26,
         1,
         1,
     }},
    {69968, 70003, 26, {}},
    {70006, 70007, 26, {}},
    {70019, 70067, 26, {}},
    {70081, 70085, 26, {}},
    {70106,
     70109,
     26,
     {
         26,
         1,
     }},
    {70144, 70162, 26, {}},
    {70163, 70188, 26, {}},
    {70207, 70209, 26, {}},
    {70272,
     70302,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         1,
         26,
         26,
         26,
         26,
         1,
     }},
    {70303, 70313, 26, {}},
    {70320, 70367, 26, {}},
    {70405, 70413, 26, {}},
    {70415,
     70441,
     26,
     {
         26,
         26,
         1,
         1,
     }},
    {70442,
     70462,
     26,
     {
         26, 26, 26, 26, 26, 26, 26, 1, 26, 26, 1, 26, 26, 26, 26, 26, 1, 1, 1,
     }},
    {70480, 70481, 26, {}},
    {70493, 70498, 26, {}},
    {70656, 70709, 26, {}},
    {70727, 70731, 26, {}},
    {70751, 70754, 26, {}},
    {70784, 70832, 26, {}},
    {70852,
     70856,
     26,
     {
         26,
         26,
         1,
     }},
    {71040, 71087, 26, {}},
    {71128, 71132, 26, {}},
    {71168, 71216, 26, {}},
    {71236, 71237, 26, {}},
    {71296, 71339, 26, {}},
    {71352, 71353, 26, {}},
    {71424, 71451, 26, {}},
    {71488, 71495, 26, {}},
    {71680, 71724, 26, {}},
    {71840, 71904, 26, {}},
    {71935, 71943, 26, {}},
    {71945,
     71956,
     26,
     {
         26,
         1,
         1,
     }},
    {71957,
     71984,
     26,
     {
         26,
         26,
         1,
     }},
    {71999,
     72002,
     26,
     {
         26,
         1,
     }},
    {72096, 72104, 26, {}},
    {72106, 72145, 26, {}},
    {72161,
     72164,
     26,
     {
         26,
         1,
     }},
    {72192, 72193, 26, {}},
    {72203, 72243, 26, {}},
    {72250, 72251, 26, {}},
    {72272, 72273, 26, {}},
    {72284, 72330, 26, {}},
    {72349, 72350, 26, {}},
    {72368, 72441, 26, {}},
    {72704, 72713, 26, {}},
    {72714, 72751, 26, {}},
    {72768, 72769, 26, {}},
    {72818, 72848, 26, {}},
    {72960,
     73009,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {73030, 73031, 26, {}},
    {73056,
     73098,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {73112, 73113, 26, {}},
    {73440, 73459, 26, {}},
    {73474,
     73489,
     26,
     {
         26,
         1,
     }},
    {73490, 73524, 26, {}},
    {73648, 73649, 26, {}},
    {73728, 74650, 26, {}},
    {74880, 75076, 26, {}},
    {77712, 77809, 26, {}},
    {77824, 78896, 26, {}},
    {78913, 78919, 26, {}},
    {82944, 83527, 26, {}},
    {92160, 92729, 26, {}},
    {92736, 92767, 26, {}},
    {92784, 92863, 26, {}},
    {92880, 92910, 26, {}},
    {92928, 92976, 26, {}},
    {92992, 92996, 26, {}},
    {93027, 93048, 26, {}},
    {93053, 93072, 26, {}},
    {93760, 93824, 26, {}},
    {93952, 94027, 26, {}},
    {94032, 94033, 26, {}},
    {94099, 94112, 26, {}},
    {94176,
     94180,
     26,
     {
         26,
         26,
         1,
     }},
    {94208, 100344, 26, {}},
    {100352, 101590, 26, {}},
    {101632, 101641, 26, {}},
    {110576,
     110883,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {110898, 110899, 26, {}},
    {110928,
     110934,
     26,
     {
         26,
         26,
         26,
         1,
         1,
     }},
    {110948, 110952, 26, {}},
    {110960, 111356, 26, {}},
    {113664, 113771, 26, {}},
    {113776, 113789, 26, {}},
    {113792, 113801, 26, {}},
    {113808, 113818, 26, {}},
    {119808, 119893, 26, {}},
    {119894, 119965, 26, {}},
    {119966,
     119994,
     26,
     {
         26,
         26,
         1,
         1,
         26,
         1,
         1,
         26,
         26,
         1,
         1,
         26,
         26,
         26,
         26,
         1,
     }},
    {119995,
     120070,
     26,
     {
         26,
         1,
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {120071,
     120085,
     26,
     {
         26,
         26,
         26,
         26,
         1,
         1,
     }},
    {120086,
     120122,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {120123,
     120486,
     26,
     {
         26, 26, 26, 26, 1,  26, 26, 26, 26, 26, 1, 26,
         1,  1,  1,  26, 26, 26, 26, 26, 26, 26, 1,
     }},
    {120488, 120513, 26, {}},
    {120514, 120539, 26, {}},
    {120540, 120571, 26, {}},
    {120572, 120597, 26, {}},
    {120598, 120629, 26, {}},
    {120630, 120655, 26, {}},
    {120656, 120687, 26, {}},
    {120688, 120713, 26, {}},
    {120714, 120745, 26, {}},
    {120746, 120771, 26, {}},
    {120772, 120780, 26, {}},
    {122624, 122655, 26, {}},
    {122661, 122667, 26, {}},
    {122928, 122990, 26, {}},
    {123136, 123181, 26, {}},
    {123191, 123198, 26, {}},
    {123214, 123215, 26, {}},
    {123536, 123566, 26, {}},
    {123584, 123628, 26, {}},
    {124112, 124140, 26, {}},
    {124896,
     124927,
     26,
     {
         26,
         26,
         26,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         26,
         1,
         26,
         26,
         1,
     }},
    {124928, 125125, 26, {}},
    {125184, 125252, 26, {}},
    {125259, 125260, 26, {}},
    {126464,
     126496,
     26,
     {
         26,
         26,
         26,
         26,
         1,
     }},
    {126497,
     126515,
     26,
     {
         26,
         26,
         1,
         26,
         1,
         1,
         26,
         1,
     }},
    {126516,
     126602,
     26,
     {
         26, 26, 26, 26, 1,  26, 1,  26, 1,  1,  1,  1,  1,  1,  26, 1,
         1,  1,  1,  26, 1,  26, 1,  26, 1,  26, 26, 26, 1,  26, 26, 1,
         26, 1,  1,  26, 1,  26, 1,  26, 1,  26, 1,  26, 1,  26, 26, 1,
         26, 1,  1,  26, 26, 26, 26, 1,  26, 26, 26, 26, 26, 26, 26, 1,
         26, 26, 26, 26, 1,  26, 26, 26, 26, 1,  26, 1,
     }},
    {126603, 126620, 26, {}},
    {126625,
     126652,
     26,
     {
         26,
         26,
         26,
         1,
         26,
         26,
         26,
         26,
         26,
         1,
     }},
    {131072, 173792, 26, {}},
    {173824, 177978, 26, {}},
    {177984, 178206, 26, {}},
    {178208, 183970, 26, {}},
    {183984, 191457, 26, {}},
    {194560, 195102, 26, {}},
    {196608, 201547, 26, {}},
    {201552, 205744, 26, {}},
};

uint32_t mapRune(int32_t c) {
  uint32_t lo = 0;
  uint32_t hi = tmRuneRanges.size();
  while (lo < hi) {
    uint32_t m = lo + (hi - lo) / 2;
    const MapRange& r = tmRuneRanges[m];
    if (c < r.lo) {
      hi = m;
    } else if (c >= r.hi) {
      lo = m + 1;
    } else {
      uint32_t i = c - r.lo;
      if (i < r.val.size()) {
        return static_cast<uint32_t>(r.val[i]);
      }
      return static_cast<uint32_t>(r.default_val);
    }
  }
  return 1;
}

constexpr int tmStateMap[] = {
    0,
};

constexpr int8_t tmToken[] = {
    1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 16,
};

constexpr int8_t tmLexerAction[] = {
    -5,  -4,  30,  -4,  23,  -4,  -4,  22,  21,  -4,  17,  16,  10,  9,   6,
    5,   5,   5,   4,   -4,  3,   5,   5,   5,   2,   1,   -4,  -7,  -7,  -7,
    -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,
    -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -7,  -6,  -6,  -6,  -6,  -6,  -6,
    -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,  -6,
    -6,  -6,  -6,  -6,  -6,  -6,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,
    -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,  -9,
    -9,  -9,  -9,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,
    -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,  -8,
    -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, 5,   5,   -16, 5,
    5,   5,   5,   -16, -16, -16, 5,   5,   5,   -16, -16, -16, -16, -16, -16,
    -1,  -16, -16, -16, -16, -16, -16, -16, 5,   5,   -16, 5,   5,   5,   5,
    -16, -16, -16, 5,   5,   5,   -16, -16, -16, -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  8,   8,   8,   8,   -4,  -4,  -4,
    8,   8,   8,   -4,  -4,  8,   -22, -22, -22, -22, -22, -22, -22, -22, -22,
    -22, -22, -22, -22, -22, 8,   8,   8,   8,   -22, -22, -22, 8,   8,   8,
    -22, -22, 8,   -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
    -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
    -15, -15, -15, -15, -15, -15, -15, -15, -15, -2,  -15, 10,  10,  -15, -15,
    -15, -3,  -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -4,  -4,  -4,
    -4,  -4,  -4,  13,  -4,  13,  -4,  -4,  12,  12,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -15, -15, -15, -15, -15, -15,
    -15, -15, -15, -15, -15, 12,  12,  -15, -15, -15, -15, -15, -15, -15, -15,
    -15, -15, -15, -15, -15, -15, -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  12,  12,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  15,
    15,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, 15,  15,  -15, -15,
    -15, -3,  -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
    -15, -15, -15, -15, -15, -15, -2,  -15, -15, -15, -15, -15, -15, -3,  -15,
    -15, -15, -15, -15, -15, -15, -15, -15, -15, -4,  -4,  -4,  -4,  -4,  18,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  18,  18,  18,  18,  19,  18,  18,  18,
    18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,
    18,  18,  18,  -4,  18,  18,  18,  18,  19,  18,  18,  18,  18,  20,  18,
    18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,  18,
    -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
    -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  16,  10,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -11, -11, -11, -11, -11, -11,
    -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
    -11, -11, -11, -11, -11, -11, -4,  23,  23,  23,  29,  23,  23,  23,  23,
    23,  23,  23,  23,  23,  23,  23,  23,  23,  23,  24,  23,  23,  23,  23,
    23,  23,  23,  -4,  -4,  -4,  -4,  23,  -4,  -4,  -4,  -4,  -4,  23,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  23,  -4,  23,  23,  25,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  26,  26,  -4,  26,
    26,  26,  -4,  -4,  -4,  -4,  26,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  27,  27,  -4,  27,  27,  27,  -4,
    -4,  -4,  -4,  27,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  -4,  -4,  -4,  28,  28,  -4,  28,  28,  28,  -4,  -4,  -4,  -4,
    28,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,  -4,
    -4,  -4,  23,  23,  -4,  23,  23,  23,  -4,  -4,  -4,  -4,  23,  -4,  -4,
    -4,  -4,  -4,  -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
    -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
    -12, -12, 30,  -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
    -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
};

constexpr int tmBacktracking[] = {
    12, 7,   // in id
    11, 14,  // in JSONNumber
    11, 11,  // in JSONNumber
};
}  // namespace

constexpr uint32_t runeErr = 0xfffd;

inline uint32_t decodeRune(std::string_view input, int64_t& offset) {
  uint8_t b0 = input[offset++];
  if (b0 < 0x80) {
    return b0;  // ASCII
  }
  uint8_t head = (b0 >> 3) & 0xf;
  int sz =
      (0x3a550000 >> (head * 2)) & 3;  // 0b10xx -> 1, 0b110x -> 2, 0b1110 -> 3
  if (sz == 0 || offset + sz > input.size()) {
    return runeErr;
  }
  uint8_t b1 = input[offset++];
  if (b1 < 0x80 || b1 >= 0xc0) {
    return runeErr;
  }
  if (sz == 1) {
    return (static_cast<uint32_t>(b0) & 0x1F) << 6 | (b1 & 0x3F);
  }
  uint8_t b2 = input[offset++];
  if (b2 < 0x80 || b2 >= 0xc0) {
    return runeErr;
  }
  if (sz == 2) {
    return (static_cast<uint32_t>(b0) & 0xF) << 12 |
           (static_cast<uint32_t>(b1) & 0x3F) << 6 | (b2 & 0x3F);
  }
  uint8_t b3 = input[offset++];
  if (b3 < 0x80 || b3 >= 0xc0) {
    return runeErr;
  }
  return (static_cast<uint32_t>(b0) & 0x7) << 18 |
         (static_cast<uint32_t>(b1) & 0x3F) << 12 |
         (static_cast<uint32_t>(b2) & 0x3F) << 6 | (b3 & 0x3F);
}

Lexer::Lexer(absl::string_view input_source) {
  source_ = input_source;
  if (absl::StartsWith(source_, bomSeq)) {
    offset_ += bomSeq.size();
  }
  Rewind(offset_);
}

Token Lexer::Next() {
restart:
  token_line_ = line_;
  token_column_ = offset_ - line_offset_ + 1;
  token_offset_ = offset_;

  int state = tmStateMap[start_state_];
  uint32_t hash = 0;
  int backupRule = -1;
  uint64_t backupOffset;
  uint32_t backupHash = hash;
  while (state >= 0) {
    int curr_class;
    if (input_rune_ < 0) {
      state = tmLexerAction[state * tmNumClasses];
      if (state > tmFirstRule && state < 0) {
        state = (-1 - state) * 2;
        backupRule = tmBacktracking[state];
        backupOffset = offset_;
        backupHash = hash;
        state = tmBacktracking[state + 1];
      }
      continue;
    } else if (input_rune_ < tmRuneClassLen) {
      curr_class = tmRuneClass[input_rune_];
    } else {
      curr_class = mapRune(input_rune_);
    }
    state = tmLexerAction[state * tmNumClasses + curr_class];
    if (state > tmFirstRule) {
      if (state < 0) {
        state = (-1 - state) * 2;
        backupRule = tmBacktracking[state];
        backupOffset = offset_;
        backupHash = hash;
        state = tmBacktracking[state + 1];
      }
      hash = hash * 31 + static_cast<uint32_t>(input_rune_);
      if (input_rune_ == '\n') {
        line_++;
        line_offset_ = offset_;
      }

      // Scan the next character.
      offset_ = scan_offset_;
      if (offset_ < source_.size()) {
        input_rune_ = decodeRune(source_, scan_offset_);
      } else {
        input_rune_ = -1;
      }
    }
  }

  int rule = tmFirstRule - state;
recovered:
  switch (rule) {
    case 12:
      switch (hash & 7) {
        case 1:
          if (hash == 0x41 && "A" == Text()) {
            rule = 16;
            break;
          }
          break;
        case 2:
          if (hash == 0x42 && "B" == Text()) {
            rule = 17;
            break;
          }
          break;
        case 3:
          if (hash == 0x5cb1923 && "false" == Text()) {
            rule = 15;
            break;
          }
          break;
        case 6:
          if (hash == 0x36758e && "true" == Text()) {
            rule = 14;
            break;
          }
          break;
        case 7:
          if (hash == 0x33c587 && "null" == Text()) {
            rule = 13;
            break;
          }
          break;
      }
      break;
    default:
      break;
  }

  Token tok = Token(tmToken[rule]);
  bool space = false;
  switch (rule) {
    case 0:
      if (backupRule >= 0) {
        rule = backupRule;
        hash = backupHash;
        Rewind(backupOffset);
      } else if (offset_ == token_offset_) {
        Rewind(scan_offset_);
      }
      if (rule != 0) {
        goto recovered;
      }
      break;
    case 8:  // space: /[\t\r\n ]+/
      space = true;
      break;
    case 18:  // 'A': /A!\p{L}+/
    {         /*some code */
    } break;
  }
  if (space) {
    goto restart;
  }
  return tok;
}

void Lexer::Rewind(int64_t rewind_offset) {
  if (rewind_offset < offset_) {
    for (int64_t i = rewind_offset; i < offset_; ++i) {
      if (source_[i] == '\n') {
        line_--;
      }
    }
  } else {
    if (rewind_offset > source_.size()) {
      rewind_offset = source_.size();
    }
    for (int64_t i = offset_; i < rewind_offset; ++i) {
      if (source_[i] == '\n') {
        line_++;
      }
    }
  }
  // Looking for \n before and not at offset_.
  line_offset_ = 1 + source_.find_last_of('\n', offset_ - 1);

  // Scan the next character.
  scan_offset_ = rewind_offset;
  offset_ = rewind_offset;
  if (offset_ < source_.size()) {
    input_rune_ = decodeRune(source_, scan_offset_);
  } else {
    input_rune_ = -1;  // Invalid rune for end of input
  }
}

}  // namespace json

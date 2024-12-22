import copy
#
# wrong btw
# def minDiffDict(myDict: dict):
#     # e1 = ""
#     # e2 = ""
#     diff = math.inf
#     prev = None
#     for ks in myDict:
#         print("outer")
#         prev = myDict[ks]
#         for k in myDict:
#             if ks == k:
#                 continue
#             newDiff = abs(prev - myDict[k])
#             if newDiff < diff:
#                 diff = newDiff
#                 print(diff)
#
#
# mDict = {
#     "joe": 29,
#     "oklj": 5,
#     "madlas": 21,
#     "saif": 3,
#     "dhflssad": 3424,
#     "323": 3423,
#     "salah": 40,
#     "shdsf": 6,
#     "dfhlasd": 2,
# }
# minDiffDict(mDict)


# def minDiffArray(arr: list[int]):
#     arr.sort()
#     minDiff = math.inf
#     newArr = []
#     for i in range(1, len(arr)):
#         temp = arr[i] - arr[i - 1]
#         if temp <= minDiff:
#             minDiff = temp
#             newArr.append([arr[i - 1], arr[i]])
#         if minDiff < (newArr[0][1] - newArr[0][0]):
#             newArr.pop(0)
#     return arr
#
#
# myArr = [-17, 46, 63, 81, -101, -91, 121, -2, 112, -15, -65, -96, 6, -139]
# minDiffArray(myArr)


def kidsWithCandies(candies: list[int], extraCandies: int):
    nkids = len(candies)
    res = []
    new = copy.deepcopy(candies)
    new.sort()
    max = new[-1]

    for i in range(nkids):
        if (candies[i] + extraCandies) >= max:
            res.append(True)
        else:
            res.append(False)
    print(max, res)


myArr = [2, 3, 5, 1, 3]
kidsWithCandies(myArr, 3)

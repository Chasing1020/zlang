let factory = function (i) {
    if (i == 0) {
        return 1;
    } else {
        return i * factory(i - 1);
    }
}

println(factory(5))

let student = {
    "name": "zjc",
    "age": 20,
}

println(student["age"])
println(student["name"])

let a = ["1", function (a, b) {
    return a + b;
}]
println(a[1](1, 3))
println(a[0])
println(len(a))

let nums = newArray(10);
println(nums);

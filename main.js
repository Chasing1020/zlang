let factory = function (i) {
    if (i == 0) {
        return 1;
    } else {
        return i * factory(i - 1);
    }
}

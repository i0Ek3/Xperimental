function insertionSort(uint[] memory a) public pure returns(uint[] memory) {
    for (uint i = 1; i < a.length; i++){
        uint tmp = a[i];
        uint j = i;
        // uint cannot be negative value, so j must >= 1
        while((j >= 1) && (tmp < a[j-1])){
            a[j] = a[j-1];
            j--;
        }
        a[j] = tmp;
    }
    return(a);
}

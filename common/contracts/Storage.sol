// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

contract Storage {
    event Stored(string indexed key, string value);

    mapping(string => string) public collection;

    function store(string memory _key, string memory _value) public {
        collection[_key] = _value;
        emit Stored(_key, _value);
    }
}

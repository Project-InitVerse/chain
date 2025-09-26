[33mcommit 8652b82817426a2b626df6c5696edaf122f15ad2[m[33m ([m[1;36mHEAD[m[33m -> [m[1;32mprivate_branch[m[33m)[m
Author: lukadas12345 <lukadas12345>
Date:   Fri Sep 12 16:32:07 2025 +0800

    add consensue ban list

[33mcommit b1f9d117c0662a25713951d3135891bc4e75a648[m[33m ([m[1;31morigin/private_branch[m[33m, [m[1;31morigin/HEAD[m[33m)[m
Author: lukads12345 <lukads12345>
Date:   Fri Aug 29 14:38:35 2025 +0800

    core/consensus: 1. add setblacklist

[33mcommit 6d7735afac935bdde7644fabdb401d565d7e00a1[m
Author: lukads12345 <lukads12345>
Date:   Thu Aug 28 15:03:47 2025 +0800

    core/consensus: 1. add setblacklist

[33mcommit 79c6a52e8c380b6285818a48bec69b3463a01036[m
Merge: 48cf61f 79e444b
Author: lukads12345 <lukads12345>
Date:   Tue Aug 5 14:10:53 2025 +0800

    Merge branch 'dev' into private_branch
    
    # Conflicts:
    #       .github/workflows/build.yml

[33mcommit 79e444be68a679b937a1b29df35d0a6f8631d941[m
Author: lukads12345 <lukads12345>
Date:   Tue Aug 5 09:33:09 2025 +0800

    core/consensus: 1. Added main chain 564480 fork code

[33mcommit 48cf61f2ae276ec6f60b881b8806f2b0d5298139[m
Merge: 5114b5e c52de62
Author: lukads12345 <lukads12345>
Date:   Mon Jul 28 11:00:02 2025 +0800

    Merge branch 'dev' into private_branch
    
    # Conflicts:
    #       consensus/inihash/inihash.go

[33mcommit 5114b5e75b0b33ce4e900e519e8ae3ab65221c82[m
Author: lukads12345 <lukads12345>
Date:   Fri Jul 25 17:24:12 2025 +0800

    fixed build deps

[33mcommit a01ec96d37776982c9316da077ae77c677a2d4c4[m
Merge: f43f6b2 c207afc
Author: lukads12345 <lukads12345>
Date:   Fri Jul 25 16:33:41 2025 +0800

    Merge branch 'main' into private_branch
    
    # Conflicts:
    #       .github/workflows/build.yml
    #       consensus/inihash/consensus.go
    #       consensus/inihash/inihash.go
    #       eth/ethconfig/config.go
    #       internal/debug/flags.go

[33mcommit 6d074a570b683e6e487e57d6fbdf137ab1d769b0[m
Author: SullivanPrime <sullivanchen93@gmail.com>
Date:   Mon Jul 21 10:02:57 2025 +0000

    add dev workflow

[33mcommit c52de62771cf2738c79981e528c3e5d633aa9700[m
Author: lukads12345 <lukads12345>
Date:   Mon Jul 21 15:44:21 2025 +0800

    consensus/inihash: Test accelerated block generation,fix reward multiplier

[33mcommit 7b59c7718e55d19d2658403007d4d72ed759f24a[m
Author: lukads12345 <lukads12345>
Date:   Mon Jul 21 11:55:57 2025 +0800

    consensus/inihash: Test accelerated block generation

[33mcommit c207afc7f0c44175970244a8a0ee02c545752c40[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 16 17:57:57 2025 +0800

    internal/debug: remove memsize (#16)
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit d4ff884f05667a6432ab9b3db7798a408849bcd8[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 16 16:36:32 2025 +0800

    core: use less memory during reorgs (#15)
    
    This PR significantly reduces the memory consumption of a long reorg
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit f43f6b297f210857621cd3581774e24f8aff81b1[m
Merge: 87e958b 8110abb
Author: lukads12345 <lukads12345>
Date:   Mon Jun 16 11:07:18 2025 +0800

    Merge branch 'private_branch' of https://github.com/Project-InitVerse/private-chain into private_branch

[33mcommit 87e958beb48df79c6aef27b08685b9c7b6134ba3[m
Author: lukads12345 <lukads12345>
Date:   Mon Jun 16 11:04:51 2025 +0800

    fix minerstartheight force set bug

[33mcommit 29baa11b95f1a41a7ab2743f7d6cdb8d1f7d1c68[m
Author: joaquinps <71144144+joaquinps@users.noreply.github.com>
Date:   Mon Jun 2 13:43:49 2025 +0800

    fix: update branch references from master and genesis-testnet to main in build.yml (#13)

[33mcommit b2ed639242bd06ca9155ffc726db25211206164d[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 2 13:27:03 2025 +0800

    accounts/abi: return error on fixed bytes with size larger than 32 bytes (#10)

[33mcommit 2ff97b93fe1d948c18f4c19342041f6cd3b72a92[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 2 13:26:26 2025 +0800

    p2p: use errors.Is for error comparison (#11)

[33mcommit a9b69e23417dfd4b2202ecf8238631550483505d[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 2 13:24:21 2025 +0800

    accounts/abi: add basic support for error types (#12)

[33mcommit 53656162c94a664261287d11af3bec0609a2a934[m
Author: joaquinps <sullivan199311@outlook.com>
Date:   Mon Jun 2 13:20:53 2025 +0800

    refactor: update import paths from PureChain to Project-InitVerse

[33mcommit 8110abbc0420e36675a65829ffd57316cdb7ea52[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Tue Apr 22 07:29:24 2025 +0000

    change workflow

[33mcommit 1b375e258b022661ed980e12f3314f033a70ca19[m
Author: lukads12345 <lukads12345>
Date:   Tue Apr 22 15:02:20 2025 +0800

    fix bugs

[33mcommit bd9a815e7838c9f87b6bdb3bb868b9e4c96961a5[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Apr 9 15:21:35 2025 +0800

    Change go mod (#9)
    
    * common/types: add Address.Big
    Many of the other types have a function to convert the type to a big.Int,
    but Address was missing this function.
    
    It is useful to be able to turn an Address into a big.Int when doing
    EVM-like computations natively in Go. Sometimes a Solidity address
    type is casted to a uint256 and having a Big method on the Address
    type makes this easy.
    
    * all: Modify the package name
    
    * cmd/clef: list accounts at startup
    Reports accounts known to Clef during startup, after master seed is provided by the user.
    
    ---------
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit a0af51267d0a061462b7492b8d6ac0712ae6f3c5[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Fri Mar 21 16:07:38 2025 +0800

    common/types: add Address.Big (#7)
    
    Many of the other types have a function to convert the type to a big.Int,
    but Address was missing this function.
    
    It is useful to be able to turn an Address into a big.Int when doing
    EVM-like computations natively in Go. Sometimes a Solidity address
    type is casted to a uint256 and having a Big method on the Address
    type makes this easy.
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit 71bcf690d99b73d54f0e2e97e0af378283acdd3f[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Mar 17 08:35:09 2025 +0800

    Add query timeout (#5)
    
    * inihash/consensus: fix incorrect comments.
    remove unused function.
    core/consensus: fix incorrect uncle block error msg
    inihash/api: remove unused comments
    
    * graphql: add query timeout
    This PR adds a 60 second timeout to graphql queries.
    
    ---------
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit 104c5231c2153fd03e3fa023ca3048b9c05d63f4[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Sat Mar 15 21:27:09 2025 +0800

    doc: update audit security report (#6)

[33mcommit ff51411e3a3ec9512b5e69675a26c78431477919[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Mar 5 16:23:35 2025 +0800

    inihash/consensus: fix incorrect comments. (#4)
    
    remove unused function.
    core/consensus: fix incorrect uncle block error msg
    inihash/api: remove unused comments
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit 137d2e2b8a527de82467c428c1ff5e81ddec7c1c[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Mar 4 14:55:24 2025 +0800

    internal/guide: speed up tests a bit (#3)
    
    updates the
    internal/guide tests to use lighter scrypt parameters.
    
    Co-authored-by: lukads12345 <lukads12345>

[33mcommit b2d3ff8425f2b8bab0e4790f516033732d7443e1[m
Author: lukads12345 <lukads12345>
Date:   Mon Feb 24 16:33:42 2025 +0800

    internal/web3ext: fix eth_call stateOverrides in console

[33mcommit c36140165fc92ab7dc3683a8036c638336709e69[m
Author: lukads12345 <lukads12345>
Date:   Wed Feb 19 13:58:46 2025 +0800

    log: use atomic type
    Modify atomic usage

[33mcommit d544fcddcc400a4afd787fa3285bb6e662a0971e[m
Author: lukads12345 <lukads12345>
Date:   Fri Feb 14 16:00:54 2025 +0800

    log: improve documentation
    Add usage examples

[33mcommit a29410c41d8cd3b39a2debf024ed685bee697871[m
Author: lukads12345 <lukads12345>
Date:   Thu Feb 6 16:32:58 2025 +0800

    ## Release Note
    
    1. change default http write timeout
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit 747f61d70f63d278418fb4e9802e50dceb211082[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Sat Jan 18 05:59:48 2025 +0000

    :sparkles: feat: add testnet bootnode

[33mcommit 51d0bc614fbde783398eca8d3ec074bea0aed97b[m
Author: lukads12345 <lukads12345>
Date:   Sat Jan 11 11:24:01 2025 +0800

    ## Release Note
    1. Fixed a pause in coin hitting caused when a transaction from a faucet was lost during a block rollback
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit e764448ce165f997776bfcde8ee7f095008ec774[m
Author: lukads12345 <lukads12345>
Date:   Sat Jan 11 00:50:03 2025 +0800

    ## Release Note
    1. Fixed a pause in coin hitting caused when a transaction from a faucet was lost during a block rollback
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit 7207cce88c6e1ab1814b92abcfb84b4adee4a03c[m
Merge: b442d53 2e1a320
Author: lukads12345 <lukads12345>
Date:   Wed Jan 8 15:24:39 2025 +0800

    Merge branch 'genesis-testnet' of https://github.com/Project-UbiCloud/ubic-chain into genesis-testnet

[33mcommit b442d53091e2c4387746157f73591016bc25ace6[m
Author: lukads12345 <lukads12345>
Date:   Wed Jan 8 15:22:32 2025 +0800

    ## Release Note
    
    1. Fixed a pause in coin hitting caused when a transaction from a faucet was lost during a block rollback
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit 2e1a3209f4cc0c38c2f199cd469098bc2d2d54e7[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Tue Jan 7 14:26:28 2025 +0800

    :zap: perf: add Initialize nodes

[33mcommit 06d66a7b7e06d65d049dd3d4bd793314156d8a0f[m
Author: lukads12345 <lukads12345>
Date:   Tue Jan 7 12:05:41 2025 +0800

    ## Release Note
    
    1. Fixed a pause in coin hitting caused when a transaction from a faucet was lost during a block rollback
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit e4487b4d4bcb7849e386cb406b0eb5d682ee917b[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Sat Jan 4 20:15:03 2025 +0800

    :memo: docs: update readme

[33mcommit ec3bbcd092409b9d8610d4b2be5e7c617fd2a3e7[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Sat Jan 4 05:07:32 2025 +0000

    :memo: docs: update readme

[33mcommit 143e48e3a82e9d825643c4212d8e647b72e7819c[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Thu Jan 2 14:53:33 2025 +0000

    :sparkles: feat: add mainnet bootnode

[33mcommit 25ae5e5e4bf40b1b793f314a9713a0e08a99cff6[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Thu Jan 2 11:34:22 2025 +0000

    :memo: docs: update readme

[33mcommit 6c4e1b63432df567e52dcb12cf361808ed4b3e0f[m
Author: SullivanPrime <sullivan199311@outlook.com>
Date:   Tue Dec 31 18:36:58 2024 +0000

    change bootnode

[33mcommit 88d9daa1e07b83aaf1ceec19aeb05af8dc6cd6e0[m
Author: lukads12345 <lukads12345>
Date:   Tue Dec 31 18:29:29 2024 +0800

    ## Release Note
    
    1. add new bootnode peer
    
    ### New Functions and Features
    
    ### New Updates and Enhancements
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit 6c0619bd6ff64d3421daf1193cf2e9c8369b0517[m
Author: lukads12345 <lukads12345>
Date:   Mon Dec 30 15:00:02 2024 +0800

    ## Release Note
    
    1. Modified configuration of chain parameters
    2. Modified some api interfaces
    
    ### New Functions and Features
    
    Updated the parameters of the Genesis block, users need to update the next version
    
    ### New Updates and Enhancements
    
    Enhanced eth.getBlockReward interface to get different block rewards on the main and test nets
    
    ### Fixed Issues
    
    ### Acknowledged Issues and Suggestions
    
    ### How to Update
    
    Recompile the client according to the latest code

[33mcommit b6aa499894e7fa194db2f704741fb54ca3ab1429[m
Author: lukads12345 <lukads12345>
Date:   Thu Dec 26 11:17:33 2024 +0800

    modify faucet page

[33mcommit 8f6f981ba54427c200f710aca5b20cb397ad65d3[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Dec 25 14:34:59 2024 +0800

    Update build.yml
    
    change workflow branch name

[33mcommit 75d137a8ae378f24d77d011b7c6846a748f06cbe[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Dec 25 13:42:03 2024 +0800

    Update faucet.go

[33mcommit 146ec998157cb9c0f5de0e144b52f601c8140ad3[m
Author: lukads12345 <lukads12345>
Date:   Mon Dec 9 15:35:01 2024 +0800

    change default testnet genesisHash

[33mcommit 44e8c8e0e29329d2cea33414c927d62ff532f93b[m
Merge: 38f14dc 8b5e5d5
Author: lukads12345 <lukads12345>
Date:   Mon Dec 9 15:15:01 2024 +0800

    tMerge branch '3.0' of https://github.com/Project-UbiCloud/ubic-chain into 3.0

[33mcommit 38f14dc29ac4ca1454d869cce8816ce7d1cc4ecc[m
Author: lukads12345 <lukads12345>
Date:   Mon Dec 9 14:48:56 2024 +0800

    modify testnet params

[33mcommit 8b5e5d5dd0496172877d6050a665065e027839ee[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Dec 9 10:17:54 2024 +0800

    Update build.yml

[33mcommit 2c3b60e0f589937c852861bc6a9ec67ee9e140a6[m
Author: lukads12345 <lukads12345>
Date:   Mon Dec 9 09:49:24 2024 +0800

    modify testnet params

[33mcommit c1350c097394be97617050e9a915fdf60ab27952[m
Author: lukads12345 <lukads12345>
Date:   Thu Nov 28 10:48:38 2024 +0800

    remove team reward

[33mcommit 3b3d2306159eef723131c190693ecac874bc7931[m
Author: lukads12345 <lukads12345>
Date:   Thu Nov 28 10:44:09 2024 +0800

    remove team reward

[33mcommit c31977461b564925f3418ca89901bd1828c55e84[m
Author: lukads12345 <lukads12345>
Date:   Wed Nov 27 12:54:10 2024 +0800

    change hash to fortihash

[33mcommit ee540e52e7ccb80f3c219d54b75603f5753a8ff4[m
Author: lukads12345 <lukads12345>
Date:   Wed Nov 20 11:42:38 2024 +0800

    change hash to fortihash

[33mcommit 8dfa949ca52bac5ac81883102ba87f4378b62b88[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 25 14:40:22 2024 +0800

    fix bugs

[33mcommit a282d7ee97fd473767b0290ce6562ddb223d9393[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 16:11:27 2024 +0800

    fix bugs

[33mcommit 8b419f93daeaa234eff2f310ff083bca6addab00[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 16:05:20 2024 +0800

    fix bugs

[33mcommit 0402fe6c6e8fc13ee8fe1bded5cd9fde8a9be180[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 15:54:18 2024 +0800

    fix bugs

[33mcommit 1f36f7eda82147b1f41def9a5f3b897f6ac4cb10[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:35:59 2024 +0800

    fix bugs

[33mcommit eb11a2727ab6ad353bd3c4c5fa2ba3e0fc41116f[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:34:57 2024 +0800

    fix bugs

[33mcommit 200db242dff43832c170675a50bd53ac170efbd0[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:33:04 2024 +0800

    fix bugs

[33mcommit cf778f01521044ad867ade57d839e590d5e6f8c0[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:25:42 2024 +0800

    fix bugs

[33mcommit fff21539213eb8431e7231cedb5a709903bcd852[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:02:38 2024 +0800

    fix bugs

[33mcommit b68950cc0bc7d7e537b40c2960f95c0d6a6200bf[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 14:00:50 2024 +0800

    fix bugs

[33mcommit 109f06bcb8d0234e6a7d1be9bee87098a8bd51ca[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 13:51:55 2024 +0800

    add log

[33mcommit 61b79a8e004f9349a6102df516dbbad8e0681020[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 13:49:02 2024 +0800

    add log

[33mcommit 5128a77ba2601720e15bb98ee0645071ea6e11a7[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 12:54:48 2024 +0800

    fix data

[33mcommit 8e20313b40c43b6d4f8b83ef478dc7e484931f28[m
Author: lukads12345 <lukads12345>
Date:   Wed Sep 18 11:32:00 2024 +0800

    fix data

[33mcommit 644d81bebf9665a5f6b896c6ea5d78605018e40b[m
Author: lukads12345 <lukads12345>
Date:   Fri Jul 19 09:20:47 2024 +0800

    reduce system save gas

[33mcommit ea1f60447290d0e2a1d1ce479283f76f121d3c19[m
Author: lukads12345 <lukads12345>
Date:   Tue Jul 16 10:35:56 2024 +0800

    try fix state at block crash bug

[33mcommit b85dad3521e0ac9498a0e6682d10211979b7661e[m
Author: lukads12345 <lukads12345>
Date:   Fri Jul 5 21:28:38 2024 +0800

    remove unused logs

[33mcommit 890e4852ac57464576785e5c3eb04a017cc88cb0[m
Author: lukads12345 <lukads12345>
Date:   Thu Jul 4 15:11:14 2024 +0800

    try fix prefetcher state root bugs

[33mcommit 94e1ad0ddcf450fa7c0847760ec98231823e079e[m
Author: lukads12345 <lukads12345>
Date:   Thu Jul 4 15:07:11 2024 +0800

    try fix prefetcher state root bugs

[33mcommit d88f506de2951bbfc82792d30e06c85d11109dc3[m
Author: tedmosby <530623363@qq.com>
Date:   Thu Jul 4 14:17:46 2024 +0800

    fix lock balance

[33mcommit cfe3181e6d8a95be782bf249be20aebe443144e6[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 21:40:23 2024 +0800

    try fix prefetcher state root bugs

[33mcommit dad31c241c61429cf74baca169893f31681367f8[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 20:52:31 2024 +0800

    try fix prefetcher state root bugs

[33mcommit fcd646a9f30fed3a7ecf63cbe18674630d12e56d[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 20:35:10 2024 +0800

    try fix prefetcher state root bugs

[33mcommit bd6161c9ba656b5ed4bae94a227402ea6c11e73b[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 20:34:51 2024 +0800

    try fix prefetcher state root bugs

[33mcommit 62d403168c0a1889c0f30e0fb800ce2d282f5c43[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 19:56:43 2024 +0800

    try fix prefetcher state root bugs

[33mcommit 06e50a1aa2ef2936976ccda9dede22b6e2f8ff57[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 19:37:09 2024 +0800

    try fix prefetcher state root bugs

[33mcommit dd54352497cc1da591934b09ba281342953a773a[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 19:35:46 2024 +0800

    try fix prefetcher state root bugs

[33mcommit 1c9c5e3864ac32808ca2b278ad768fc95678d1e0[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 18:39:50 2024 +0800

    try fix prefetcher state root bugs

[33mcommit 033d05f2463641827d6e6a70b4a8a2d04c4fcfcd[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 17:04:40 2024 +0800

    add applyTransaction

[33mcommit a14a816857aa5ec28b4b8aba588dd777ff7ac763[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 12:41:58 2024 +0800

    add transaction commit root hash print

[33mcommit dbb3cb49581f6bbb2fd578b0d3009d005208b9cd[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 11:46:10 2024 +0800

    add transaction commit root hash print

[33mcommit 34ddfbc8316718fb3bfe146c35d52f66f32efc2c[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 11:18:55 2024 +0800

    add transaction commit root hash print

[33mcommit 0e26395119cac7cee668075baf1079ea20bb2855[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 10:59:47 2024 +0800

    add transaction commit root hash print

[33mcommit 126d9c04709a62302aaec82783e2329e60a9a21c[m
Author: lukads12345 <lukads12345>
Date:   Wed Jul 3 09:13:02 2024 +0800

    add transaction commit root hash print

[33mcommit 356dc8398da54d4681444fb8e3b93adfb52d0349[m
Author: lukads12345 <lukads12345>
Date:   Tue Jul 2 18:17:56 2024 +0800

    add transaction commit root hash print

[33mcommit 689542635a248277b40502e7ae2c1a5596e752cc[m
Author: lukads12345 <lukads12345>
Date:   Tue Jul 2 16:53:30 2024 +0800

    add transaction commit root hash print

[33mcommit 35281df2273e135e91d35392c9e98daff4934db4[m
Author: lukads12345 <lukads12345>
Date:   Tue Jul 2 16:27:08 2024 +0800

    try fix commit state race condition

[33mcommit c72a7e12631e3e4328690f9f116953a56290bc1f[m
Author: lukads12345 <lukads12345>
Date:   Tue Jul 2 15:19:56 2024 +0800

    try fix commit state race condition

[33mcommit d1a0c5167a18385f926a1a0c2cbacfed6cc2aed4[m
Author: lukads12345 <lukads12345>
Date:   Mon Jul 1 16:02:53 2024 +0800

    add root hash print

[33mcommit b517f0c9515ed7eaede00d147ba9d5271de655b5[m
Author: lukads12345 <lukads12345>
Date:   Mon Jul 1 15:21:18 2024 +0800

    fix database lock range

[33mcommit 21a14fbeb57a105b014c69ee9d3ca9852c9194d6[m
Author: lukads12345 <lukads12345>
Date:   Mon Jul 1 14:38:37 2024 +0800

    fix state db race condition

[33mcommit 0c34d3c07561bc933b0c341666ad9044acf933a6[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Jun 27 16:40:11 2024 +0800

    Update block_validator.go

[33mcommit 94d7a6f4420c8122c1a8fd5e53b477067e45e9e1[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Jun 27 15:00:31 2024 +0800

    Update block_validator.go

[33mcommit 72d92358e87c68a9576e6248164ceb8b603b7ed2[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Jun 27 14:53:32 2024 +0800

    Update block_validator.go

[33mcommit 97faaa999fddbd2b1f269b5285ba4c40c0b0f588[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Jun 27 14:48:45 2024 +0800

    Update block_validator.go

[33mcommit 57248c2f021abccb8626fcef3d410fc2cb9883cb[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 26 09:55:48 2024 +0800

    Update worker.go

[33mcommit 7c831dd7433d9ba1d98642026ff2039abedc7daf[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Fri Jun 21 18:34:26 2024 +0800

    Update block_validator.go

[33mcommit bbe2499c5103d5efdfc8997ad2f801f62b8bbc05[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Fri Jun 21 18:06:25 2024 +0800

    Update block_validator.go

[33mcommit 1909cc8d039adedf5ac926168c154faa46df627e[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Fri Jun 21 17:46:53 2024 +0800

    Update block_validator.go

[33mcommit b43be98cbd8a2269ff706b1c38d0e00939f8dfbf[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Fri Jun 21 17:12:26 2024 +0800

    Update block_validator.go

[33mcommit c534583e48a6e6313d8feb9721c48cb55acce54a[m
Author: lukads12345 <lukads12345>
Date:   Tue Jun 4 10:05:05 2024 +0800

    reduce block tx limit

[33mcommit 8c9c3f75d10c094cb2587c2d5ab6bb9345efd37f[m
Author: lukads12345 <lukads12345>
Date:   Tue Jun 4 09:19:36 2024 +0800

    fix panic nil pointer

[33mcommit 2864732b9fdc540743656c0b13b622af1fdf090b[m
Author: lukads12345 <lukads12345>
Date:   Mon Jun 3 23:26:12 2024 +0800

    try fix panic errors

[33mcommit bd53ca65d9b463f77667ec8e080993558139aef1[m
Author: lukads12345 <lukads12345>
Date:   Sat Jun 1 10:28:03 2024 +0800

    Reduced logging overhead

[33mcommit 6837fc6eb1412d060150d5433670396b03b08436[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Apr 25 20:19:08 2024 +0800

    Update worker.go

[33mcommit 0a503edee7e5454d1439b61fc2d5f30805681aae[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Thu Apr 25 19:36:48 2024 +0800

    Update worker.go
    
    fix in turn choose

[33mcommit bcb9a72d9a5ea36e185c151cc95cc2a641d6cd42[m
Author: snakelu <lyhccq@163.com>
Date:   Wed Feb 21 11:52:42 2024 +0800

    feat: base58 decode function return string

[33mcommit 975461e0d052edd6b7a0923a709d5bf4599ea965[m
Author: snakelu <lyhccq@163.com>
Date:   Wed Feb 21 11:36:41 2024 +0800

    feat: faucet support I4 address

[33mcommit 3a66654f9167b1141f829ef4414014a7ecc6d3e3[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Feb 19 17:05:10 2024 +0800

    feat: faucet support ipv6

[33mcommit de91bc0758b68d0a91373500022185534bc62a42[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Feb 19 15:50:46 2024 +0800

    feat: add log

[33mcommit deea51b1558d73f8456df24c785f229ab57e4811[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Sun Feb 18 16:49:54 2024 +0800

    feat: log remote addr

[33mcommit 9f3b17eea37c4aa1524b11bf632c84f495333191[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Sun Feb 18 16:35:28 2024 +0800

    fix: ip parse

[33mcommit 966ee6b533dc444bd7d97b2a0a1b62d8466f3ff3[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Sun Feb 18 16:08:14 2024 +0800

    fix: get client real ip from request

[33mcommit 6be93cd8d2899a576a30ec3308a017a490e23b12[m
Author: snakelu <lyhccq@163.com>
Date:   Tue Feb 6 21:22:56 2024 +0800

    feat: add log

[33mcommit d2f320af59d934c1ff7607cc5e1435a5ddd383ff[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Jan 15 15:53:00 2024 +0800

    feat: faucet address hint

[33mcommit 41e3b9c0a5458439992040e441b8be072515ccb7[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Sat Dec 23 13:38:02 2023 +0800

    feat: change challenge rate

[33mcommit 7025fcc61c3271064ccc559efdccb6568e67a72c[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Sat Dec 23 13:15:15 2023 +0800

    feat: change year block count

[33mcommit a9c89a3bac4a47108b6b934a17c48874e11d3fd4[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Dec 20 14:37:16 2023 +0800

    feat: testnet remove address list

[33mcommit 59e7b526570ee2fb62d5369ab1f02175a92505e8[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Wed Dec 20 10:50:40 2023 +0800

    :bug: fix: update genesis hash

[33mcommit f6c18aea3ac2701c3fc2caf3bd838ba81f6b2d99[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Dec 20 10:43:20 2023 +0800

    feat: change testnet chainid

[33mcommit 723c519aeaa19dca365f6a985d3e00a8666eb42f[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Dec 20 10:39:57 2023 +0800

    feat: testnet genesis block timestamp

[33mcommit 5e88852c8ade828c1d6701472f95ab0d936243b3[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Wed Dec 20 09:53:00 2023 +0800

    :bug: fix: error

[33mcommit 14591cb6057c3b780a84d7d7fb8726ebee621b0e[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Dec 20 09:38:55 2023 +0800

    feat: testnet genesis config

[33mcommit 1b9e31c6d1491ccafe941dd6b58e672a29bac038[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Dec 20 09:17:21 2023 +0800

    feat: testnet genesis_alloc

[33mcommit a2fbfb0ae1762f48f446f153aaa4d45ef8182a2e[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Dec 19 18:03:42 2023 +0800

    feat: testnet genesis.json

[33mcommit 1f0e848d795e23312330b961399d89b046e874fc[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Dec 19 17:29:47 2023 +0800

    feat: genesis hash

[33mcommit 457721be291cc821aaa071b22a809bd5a060092f[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Dec 19 16:40:39 2023 +0800

    feat: testnet alloc

[33mcommit 14a092675381720433ef726efa09340103a2a746[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Dec 19 15:27:40 2023 +0800

    feat: testnet alloc data

[33mcommit 2dc72cb074ea0bb2ddc7553e01b46349aeb93ba2[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Dec 19 14:22:19 2023 +0800

    feat: challenge verify fault-tolerant

[33mcommit d97bb7bec5343810975fbfb94b129e1f51cfda50[m
Author: Jerry Zhang <zherrynju@gmail.com>
Date:   Wed Nov 29 15:31:33 2023 +0800

    Update README.md

[33mcommit 6587d0443277d099f1e0a5a78144419eee762d23[m
Author: snakelu <lyhccq@163.com>
Date:   Wed Nov 29 11:28:54 2023 +0800

    Update README.md

[33mcommit 6d388c79c9a9490131e5e3ec62bba2e67ccb654a[m
Author: snakelu <lyhccq@163.com>
Date:   Wed Nov 29 11:25:14 2023 +0800

    Update README.md

[33mcommit 6a9dc35881f75157b1c04982c2656b923bad4c80[m
Author: snakelu <lyhccq@163.com>
Date:   Wed Nov 29 10:56:27 2023 +0800

    Update README.md

[33mcommit a0a14839d4683ab18d6b17bb5e3916b52a102cf4[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Nov 9 15:55:42 2023 +0800

    feat: submit seed add log

[33mcommit 27f9bda7c0778e65a3625279b1c92526b4e1b34d[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Nov 9 15:45:22 2023 +0800

    feat: add log before submit seed

[33mcommit 1b00d49f6b215f1575a0772458b8669d13aa7c73[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Oct 19 16:36:25 2023 +0800

    feat: change challenge interval

[33mcommit d43f52ce3c9b7b0c0a6af88f68a95ac4bca53cfc[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Oct 19 11:58:15 2023 +0800

    feat: update core/genesis_alloc.go

[33mcommit fbada51efa4d453b9e4f83e7e3df9c1cdd4e4df4[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Oct 18 16:35:33 2023 +0800

    fix: cannot unmarshal *big.Int in to uint64

[33mcommit 8053cd206c134757678057ab6200b77119c142e8[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Oct 18 16:06:09 2023 +0800

    feat: get provider info

[33mcommit 47ac2700a1e5b018e7eb7d786b475e43c0ebb9d3[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Oct 18 11:45:40 2023 +0800

    update genesis_alloc.go

[33mcommit 2f50e344d409788634c628f887734f6bbd4af768[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Oct 18 09:40:52 2023 +0800

    feat: refresh devnet genesis alloc

[33mcommit ac0c46255b06003c75703096cbd662d35cd3f198[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Oct 16 16:00:26 2023 +0800

    feat: refresh devnet genesis alloc

[33mcommit 1b18d73070adc1272d8d23a7b9801ccea86fc75f[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Oct 12 15:53:58 2023 +0800

    feat: modify devnet genesis

[33mcommit 4a207bf41ad861c15fda65eea03f915d1cebb120[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Tue Oct 10 15:42:57 2023 +0800

    feat: refresh providerFactory abi

[33mcommit f903f184bf431f1e33b332e4b9137562d5744a92[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Sep 25 17:33:30 2023 +0800

    feat: verify task cost mills log

[33mcommit 68b238f84628c856ab058862736393178ae8f8a6[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Sep 25 16:56:58 2023 +0800

    feat: adjust por verify leaf

[33mcommit cc8d51ab45d115be979b4f479d4dd0ae3f9acc00[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Sep 21 17:18:53 2023 +0800

    feat: query ready

[33mcommit adbfa78146b62b30c49db6cadb62698d48e7e623[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Sep 21 16:20:50 2023 +0800

    feat: change timeout

[33mcommit 0314da578bf473e7670894b06134467009d70051[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Sep 18 15:21:30 2023 +0800

    feat: add verify leaf limit & extend timeout based on challenge count

[33mcommit 8c1dd716d0dc7db5f83e8341f461dae30103e451[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Sep 4 17:17:39 2023 +0800

    fix: team & validator reawrd rate

[33mcommit 4c4fa410fbe4b96032757669538f353e75d9eafa[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Sep 4 16:51:20 2023 +0800

    fix: provider reward lock balance

[33mcommit 99f8a3eacd76b55a3766f3f8ec33b07fe8745136[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Aug 30 13:52:35 2023 +0800

    fix: verify task

[33mcommit 56ad2932647dfff1c32f3e373dd64b54879c64a3[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Thu Aug 24 17:04:52 2023 +0800

    feat: por challenge add log

[33mcommit 65d565ed384f50321f569f229d8707d1a1567e77[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Fri Aug 18 15:54:33 2023 +0800

    feat: challengeRate

[33mcommit b88121199acfda54d8bb19a0dfed918fac221f80[m
Merge: 8090b03 1f74313
Author: lukads12345 <lukads12345>
Date:   Fri Aug 18 15:44:12 2023 +0800

    Merge remote-tracking branch 'origin/2.0' into 2.0

[33mcommit 8090b03518344ef2382d718cc88ba177f440b9bc[m
Author: lukads12345 <lukads12345>
Date:   Fri Aug 18 15:44:01 2023 +0800

    fix challenge bugs

[33mcommit 1f74313791f08516ee26710b36035d83b8284bce[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Thu Aug 17 15:44:47 2023 +0800

    :sparkles: feat: add bootnodes

[33mcommit 7193914a114163bf6b52594fc0b4bb9f11209b91[m
Author: lukads12345 <lukads12345@gamil.com>
Date:   Tue Aug 15 11:13:48 2023 +0800

    add testnet & devnet admin

[33mcommit e2ddb39042c167cb82b1f77308cb9b48a7987a22[m
Author: lukads12345 <lukads12345>
Date:   Tue Aug 15 10:51:08 2023 +0800

    add --devnet flags

[33mcommit 296f36fd063aa60a50ac96cfb4d2ac69763d8a15[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Mon Aug 14 14:42:14 2023 +0800

    feat: faucet UBIC -> INI

[33mcommit 64ee45400b6877d3788170c1ffe6d965d57fae43[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Fri Aug 11 17:32:14 2023 +0800

    :sparkles: feat: generate faucet

[33mcommit 6b61e55d16331cbb3a1a62f8d7c9625ae2e62637[m
Author: xia <xiadd0102@gmail.com>
Date:   Fri Aug 11 14:38:34 2023 +0800

    Update faucet.html

[33mcommit b3724e56c04249f8219c38b2cd72b48862171c35[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Wed Aug 9 19:13:09 2023 +0800

    :sparkles: feat: set params (#1)

[33mcommit 4de7b75b62b92fe6c5a968e53910d8e3bb496da8[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Aug 9 18:15:24 2023 +0800

    feat: change U4 address prefix to I4 go-bindata

[33mcommit 6fd94c9778c03300a65171511631e080893d1204[m
Author: snakelu <luyuanheng@whitematrix.io>
Date:   Wed Aug 9 17:16:52 2023 +0800

    feat: change U4 address prefix to I4

[33mcommit e0ddaea184dacebc810dabab6ee5774fad6e88be[m
Author: lukads12345 <lukads12345@gamil.com>
Date:   Fri Jun 30 15:39:48 2023 +0800

    fix mockowner

[33mcommit 6429fb9cb1bb79d9577c844f1011d13d50468fbc[m
Author: lukads12345 <lukads12345>
Date:   Fri Jun 30 11:35:22 2023 +0800

    change faucet.html

[33mcommit eab30269db3832153febcd1d27dbe921be8d2f93[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 14:47:28 2023 +0800

    change faucet.html

[33mcommit 47bf63586d0c0a339536fe52bd962c5d85793a41[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 12:04:45 2023 +0800

    change faucet.html

[33mcommit 92d433b57d70b269c4c86c1248a6101bcac4cc8a[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 11:57:20 2023 +0800

    change faucet.html

[33mcommit df293e358dffad2b7eae3b92779e34c1529a94fe[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 11:28:42 2023 +0800

    fix get ip bugs

[33mcommit 6c4bf36ee0f77bec61a54e663fad0bd786292d4d[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 11:24:21 2023 +0800

    fix get ip bugs

[33mcommit c73b97cac8d72ae7261ee3919c019c0ba5293913[m
Merge: c0f2ffa f2492b3
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 11:06:33 2023 +0800

    Merge branch '2.0' of https://github.com/Project-UbiCloud/ubic-chain into 2.0

[33mcommit c0f2ffaa2a13f8805ee980ea1ec06db86d3d22d8[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 11:05:31 2023 +0800

    change faucet.html,fix miner flag bugs

[33mcommit f2492b37e97221e411fbe7a8796fee75d976a195[m
Author: root <root@ubicloud-master.us-east1-b.c.ml-devops-380009.internal>
Date:   Thu Jun 29 02:37:41 2023 +0000

    update genesis.json

[33mcommit 3b132e1c4f753b9cede42a296a9165c44b226f3e[m
Merge: 0714c24 65cadd5
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 10:08:33 2023 +0800

    Merge remote-tracking branch 'origin/2.0' into 2.0

[33mcommit 0714c24db899e7e44acfd39bbfc276811334049b[m
Author: lukads12345 <lukads12345>
Date:   Thu Jun 29 10:08:23 2023 +0800

    change faucet.html,fix miner flag bugs

[33mcommit 65cadd51176edb1d7d08dd55caf2daf87a920e6d[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 28 19:12:23 2023 +0800

    Update blockchain.go

[33mcommit 9b1de5274004b40a19e7515d57a8cd0e15c5cd00[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 28 19:10:08 2023 +0800

    Update blockchain.go
    
    remove unused fmt println

[33mcommit 2af12bb339176ae899937429b7ee66de47778418[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 28 19:02:09 2023 +0800

    Update state_processor.go
    
    remove fmt.print

[33mcommit d06a743adddc54e16d1e78d3d1b6f5ab001f1ba1[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 28 18:56:32 2023 +0800

    Update bootnodes.go
    
    change bootnode

[33mcommit 3f1f1963ddcc657e9b0f32077ad9f52b61978dc7[m
Author: lukads12345 <lukads12345>
Date:   Wed Jun 28 18:25:28 2023 +0800

    change faucet.html

[33mcommit 3392a40ce251cf395a4a0dab325b748b108a8250[m
Merge: e3d632a 6119a19
Author: lukads12345 <lukads12345>
Date:   Wed Jun 28 16:47:38 2023 +0800

    auto merge

[33mcommit e3d632aae17a4370375316864f08a87bac3eb7f7[m
Author: lukads12345 <lukads12345>
Date:   Wed Jun 28 16:45:38 2023 +0800

    fix pending txs cal failed

[33mcommit 6119a19d97b6175048237ff298059d73b5f893b9[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Wed Jun 28 15:13:15 2023 +0800

    :bug: fix build error

[33mcommit 596f198383e9fb1f3427c3d7942f0981dfabd9ad[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Tue Jun 27 17:52:50 2023 +0800

    :sparkles: add all tools image

[33mcommit 426f196b2692c51a2f5d0f842dbe5eca6a1af659[m
Author: lukads12345 <lukads12345>
Date:   Tue Jun 27 17:32:12 2023 +0800

    fix pending txs cal failed

[33mcommit ba23cfd68df866f65dd7d6d6a3b4622c352bc48c[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Jun 27 16:09:35 2023 +0800

    Update config.go
    
    change default sync mode from fast to full

[33mcommit 03e7b95762653418aa4255a87a7b4c6413bc239c[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Jun 27 12:18:34 2023 +0800

    Update config.go

[33mcommit 6efab0ef7219afb449167de4121c2d54b6de0cac[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Jun 27 12:05:52 2023 +0800

    Update config.go

[33mcommit 7bb5a7062f5273dce0e8384d573278370db97c6c[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Jun 27 12:01:18 2023 +0800

    Update bootnodes.go

[33mcommit e66fae4002d7695e40f1d370618a2616c9502e0d[m
Merge: ad64a32 c583167
Author: lukads12345 <lukads12345@gamil.com>
Date:   Tue Jun 27 09:12:12 2023 +0800

    Merge branch '2.0' of https://github.com/Project-UbiCloud/chain into 2.0

[33mcommit ad64a329a1b2a0b61805144e314f7994b56a15b5[m
Author: lukads12345 <lukads12345@gamil.com>
Date:   Tue Jun 27 09:10:32 2023 +0800

    add punish item contract

[33mcommit c5831675284d581e571ed3cf41b32715b9da3fae[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Mon Jun 26 09:58:59 2023 +0800

    Update worker.go
    
    fix por result transaction missed bugs

[33mcommit 9359d1680f13a5cd968f12745a06fdfdfaa2607e[m
Author: lukads12345 <lukads12345>
Date:   Mon Jun 26 09:02:48 2023 +0800

    commit README.md

[33mcommit 913dea011af509c63d04165f145644ab21f37fae[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Sun Jun 25 13:43:26 2023 +0800

    Update dpos.go
    
    fix choose provider bugs

[33mcommit 954a1ca1dd0e2719dd7f36f3250806f47b7b3bbe[m
Author: lwnmengjing <lwnmengjing@qq.com>
Date:   Thu Jun 22 12:35:26 2023 +0800

    :rocket: add workflows

[33mcommit 066c45570f54dc599e0e7ec21b4e3d24795efc3d[m
Author: lukads12345 <lukads12345>
Date:   Wed Jun 21 18:50:34 2023 +0800

    change testnet config

[33mcommit 427931d3231cb6c6574af6804a4d1c6c231568b4[m
Author: lukads12345 <lukads12345>
Date:   Wed Jun 21 18:47:16 2023 +0800

    fix worker commit challenge result failed bug

[33mcommit 877d624c2e63112159e7bc1d256cc0612bd93ef9[m
Author: lukads12345 <lukads12345>
Date:   Wed Jun 21 09:50:54 2023 +0800

    fix por challenge margin amount calculate bugs

[33mcommit 36c2a2dd66169ce0ecc27eac8ed06d0a31831193[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 14 17:29:42 2023 +0800

    Update go.mod
    
    fix go mod

[33mcommit 816eaf42def4bcfaabece7ecc1e46f167dd7e046[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 14 16:31:59 2023 +0800

    Update dpos.go
    
    add address list contract init

[33mcommit d5d0e80e21d334aac65afd4316216810b254bba5[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 14 16:29:22 2023 +0800

    Update abi.go
    
    add addresslist admin

[33mcommit 41b8eb12e2a54d6ba5f50699c88601bf882c04b9[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Wed Jun 14 16:16:06 2023 +0800

    Update types.go
    
    remove debug logs

[33mcommit 08c23ea81bcf8bb4620b7a55c6166de64c405c9b[m
Author: lukads12345 <118886886+lukads12345@users.noreply.github.com>
Date:   Tue Jun 13 18:27:10 2023 +0800

    Update mkalloc.go
    
    fix path error

[33mcommit c37ecda70e49b78d4b01483eeb82ccb6665ce52d[m
Author: lukads12345 <lukads12345>
Date:   Mon Jun 12 17:27:23 2023 +0800

    fix can challenge function bugs

[33mcommit 1905c56d3546fcd3415ce87f0ea0a554756eb164[m
Author: lukads12345 <lukads12345>
Date:   Tue May 30 11:56:09 2023 +0800

    Fixed code specification related issues

[33mcommit 477645e4eee3f7c7818b6b09fdcb0106dbf656b4[m
Author: lukads12345 <lukads12345>
Date:   Mon May 29 17:45:43 2023 +0800

    remove unused print,Fixed an issue with repeatedly creating invalid Pors

[33mcommit 80d4ac08a976b455759da054143264dc5e24b50a[m
Author: lukads12345 <lukads12345>
Date:   Mon May 29 17:23:16 2023 +0800

    The challenge scheduler code was submitted and the mining process was added to handle POR response events

[33mcommit 41f5e873f72c58a5e481a0e18a2f551821aa3d93[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Mon Apr 24 09:47:47 2023 +0800

    fix magic number

[33mcommit bd28972d8c3184cc58fc66ce53a10570d3069de2[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Mon Apr 24 09:30:26 2023 +0800

    fix magic number,commit faucet module

[33mcommit 8a50988f754b56836fa8ce288bf72b2920c19172[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Mon Apr 17 10:24:29 2023 +0800

    add eth.getLockBalance interface

[33mcommit 15111e9eb174405f0a47876ffab45d4a20f9530b[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Mon Apr 10 09:22:18 2023 +0800

    Block logic modification, token allocation, token production reduction, fee destruction

[33mcommit a7bbaf558ddec0a9965f6f26fbd29bcad5e63fc9[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Wed Feb 15 09:39:59 2023 +0800

    fix dpos distribute reward bugs.

[33mcommit a9eceae7f6f6f18905322d687bbe0d54e24f56ce[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Wed Dec 7 10:06:47 2022 +0800

    fix bugs

[33mcommit 7844757982d7aaf8aeae9c2d80d8b416e774fa65[m
Author: lukads12345 <lukads12345@gmail.com>
Date:   Tue Nov 29 14:51:13 2022 +0800

    first commit

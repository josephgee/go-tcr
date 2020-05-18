ginkgo . && (git add . && git ci -m "Passed") || (git reset --hard HEAD && git clean -ffd -e .idea)

CXX      = g++
CXXFLAGS = -std=c++17 -Wall -Wextra -O2

LR1 = array.cpp forward_list.cpp list.cpp stack.cpp queue.cpp tree.cpp

TARGETS = task1_postfix task2_set task3_partition task4_subarrays task5_bst_fix task6_routes task7_lru

all: $(TARGETS)

task1_postfix: task1_postfix.cpp $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ $^
task2_set: task2_set.cpp set.cpp $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ $^
task3_partition: task3_partition.cpp set.cpp $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ $^
task4_subarrays: task4_subarrays.cpp set.cpp $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ $^
task5_bst_fix: task5_bst_fix.cpp $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ $^
task6_routes: task6_routes.cpp hashtable.h $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ task6_routes.cpp $(LR1)
task7_lru: task7_lru.cpp hashtable.h $(LR1)
	$(CXX) $(CXXFLAGS) -o $@ task7_lru.cpp $(LR1)

clean:
	rm -f $(TARGETS)

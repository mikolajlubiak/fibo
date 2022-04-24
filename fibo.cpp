#include <vector>

void fibonacci(unsigned int x, unsigned long long int y, unsigned long long int z) {
	std::vector<unsigned long long int> fibo;
	fibo.push_back(y);
	fibo.push_back(z);
	printf("1: %llu\n2: %llu\n", fibo[0], fibo[1]);
	for (unsigned int i = 2; i < x; i++) {
		fibo.push_back(fibo[i-1] + fibo[i-2]);
		printf("%u: %llu\n", i+1, fibo[i]);
	}
}

int main(int argc, char* argv[]) {
	if (argc != 4) {
		printf("You have to provide 3 arguments - first number, secound number, how many numbers to calculate.\nExample:\t./fibo 0 1 100\n");
		return 1;
	}
	fibonacci(atoi(argv[3]), atoi(argv[1]), atoi(argv[2]));
	return 0;
}
